package motdsource

import (
	"bufio"
	"bytes"
	"context"
	"time"

	motdv1alpha1 "github.com/alanraison/motd-operator/pkg/apis/motd/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_motdsource")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new MotdSource Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &ReconcileMotdSource{client: mgr.GetClient(), scheme: mgr.GetScheme()}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("motdsource-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource MotdSource
	err = c.Watch(&source.Kind{Type: &motdv1alpha1.MotdSource{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner MotdSource
	// err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
	// 	IsController: true,
	// 	OwnerType:    &motdv1alpha1.MotdSource{},
	// })
	// if err != nil {
	// 	return err
	// }

	return nil
}

// blank assignment to verify that ReconcileMotdSource implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileMotdSource{}

// ReconcileMotdSource reconciles a MotdSource object
type ReconcileMotdSource struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme
}

// Reconcile reads that state of the cluster for a MotdSource object and makes changes based on the state read
// and what is in the MotdSource.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileMotdSource) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling MotdSource")

	// Fetch the MotdSource instance
	instance := &motdv1alpha1.MotdSource{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	motd, err := fetchMotd(instance.Spec)
	if err != nil {
		status := motdv1alpha1.MotdSourceStatus{
			Updated: metav1.Now(),
			Error:   err.Error(),
		}
		instance.Status = status
		reqLogger.Error(err, "fetching motd")
		r.client.Status().Update(context.TODO(), instance) // ignore errors
		return reconcile.Result{}, err
	}
	status := motdv1alpha1.MotdSourceStatus{
		Updated:      metav1.Now(),
		ShortMessage: shortenMotd(motd),
		FullMessage:  motd,
	}
	instance.Status = status
	if err := r.client.Status().Update(context.TODO(), instance); err != nil {
		return reconcile.Result{}, err
	}

	reqLogger.Info("Updated MotdSource; requeuing in 30s")
	return reconcile.Result{
		Requeue:      true,
		RequeueAfter: 30 * time.Second,
	}, nil
}

func shortenMotd(motd string) string {
	buf := make([]byte, 30)
	copy(buf, motd)
	scanner := bufio.NewScanner(bytes.NewReader(buf))
	scanner.Scan()
	return scanner.Text()
}

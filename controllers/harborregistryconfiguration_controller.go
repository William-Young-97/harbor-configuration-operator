/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	apiv2 "github.com/mittwald/goharbor-client/v5/apiv2"
	modelv2 "github.com/mittwald/goharbor-client/v5/apiv2/model"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	v1alpha1 "github.com/william-young-97/harbor-configuration-operator/api/v1alpha1"
)

// HarborRegistryConfigurationReconciler reconciles a HarborRegistryConfiguration object
type HarborRegistryConfigurationReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=harbor-configuration.harbor-practice.co.uk,resources=harborregistryconfigurations,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=harbor-configuration.harbor-practice.co.uk,resources=harborregistryconfigurations/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=harbor-configuration.harbor-practice.co.uk,resources=harborregistryconfigurations/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the HarborRegistryConfiguration object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.2/pkg/reconcile
func (r *HarborRegistryConfigurationReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	var harborRegistryConfiguration v1alpha1.HarborRegistryConfiguration

	err := r.Get(ctx, req.NamespacedName, &harborRegistryConfiguration)
	if err != nil {
		return ctrl.Result{}, err
	}

	client, err := apiv2.NewRESTClientForHost("https://core.harbor.harbor-practice.co.uk/api/v2.0", "admin", "Harbor12345", nil)
	if err != nil {
		return ctrl.Result{}, err
	}

	myRegistry := &modelv2.Registry{
		Name: harborRegistryConfiguration.Spec.Name,
		Type: "docker-hub",
		URL:  "https://hub.docker.com",
	}

	err = client.NewRegistry(ctx, myRegistry)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HarborRegistryConfigurationReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&v1alpha1.HarborRegistryConfiguration{}).
		Complete(r)
}

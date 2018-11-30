package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	MachineControllerFrozenDesc = prometheus.NewDesc("mcm_machine_controller_frozen", "Frozen status of the machine controller manager.", nil, nil)
	MachineCountDesc            = prometheus.NewDesc("mcm_machine_items_total", "Count of machines currently managed by the mcm.", nil, nil)
	MachineCreated              = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_created",
		Help: "Creation time of the Machines currently managed by the mcm.",
	}, []string{"name", "namespace", "uid"})

	MachineCSPhase = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_current_status_phase",
		Help: "Current status phase of the Machines currently managed by the mcm.",
	}, []string{"name", "namespace", "uid", "phase"})

	MachineInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_info",
		Help: "Information of the Machines currently managed by the mcm.",
	}, []string{"name", "namespace", "uid", "generation", "kind", "api_version",
		"spec_provider_id", "spec_class_api_group", "spec_class_kind", "spec_class_name"})

	MachineStatusCondition = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_status_condition",
		Help: "Information of the mcm managed Machines' status conditions.",
	}, []string{"name", "namespace", "uid", "condition", "status"})

	MachineSetCountDesc = prometheus.NewDesc("mcm_machineset_items_total", "Count of machinesets currently managed by the mcm.", nil, nil)
	MachineSetCreated   = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_set_created",
		Help: "Creation time of the Machinesets currently managed by the mcm.",
	}, []string{"name", "namespace", "uid"})

	MachineSetInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_set_info",
		Help: "Information of the Machinesets currently managed by the mcm.",
	}, []string{"name", "namespace", "uid", "generation", "kind", "api_version",
		"spec_machine_class_api_group", "spec_machine_class_kind", "spec_machine_class_name"})

	MachineSetInfoSpecReplicas = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_set_info_spec_replicas",
		Help: "enca countof the Machinesets currently managed by the mcm.",
	}, []string{"name", "namespace", "uid"})

	MachineSetInfoSpecMinReadySeconds = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_set_info_spec_min_ready_seconds",
		Help: "Information of the Machinesets currently managed by the mcm.",
	}, []string{"name", "namespace", "uid"})

	MachineSetStatusCondition = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_set_status_condition",
		Help: "Information of the mcm managed Machinesets' status conditions.",
	}, []string{"name", "namespace", "uid", "condition", "status"})

	MachineSetStatusFailedMachines = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_set_failed_machines",
		Help: "Information of the mcm managed Machinesets' failed machines.",
	}, []string{"name", "namespace", "uid", "failed_machine_name", "failed_machine_provider_id", "failed_machine_owner_ref",
		"failed_machine_last_operation_state",
		"failed_machine_last_operation_machine_operation_type"})

	MachineSetStatus = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_set_status",
		Help: "Information of the mcm managed Machinesets' status conditions.",
	}, []string{"name", "namespace", "uid", "available_replicas",
		"fully_labeled_replicas", "ready_replicas", "replicas"})

	MachineDeploymentCountDesc = prometheus.NewDesc("mcm_machinedeployment_items_total", "Count of machinedeployments currently managed by the mcm.", nil, nil)
	MachineDeploymentCreated   = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_deployment_created",
		Help: "Creation time of the Machinedeployments currently managed by the mcm.",
	}, []string{"name", "namespace", "uid"})
	MachineDeploymentInfo = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_deployment_info",
		Help: "Information of the Machinedeployments currently managed by the mcm.",
	}, []string{"name", "namespace", "uid", "generation", "kind", "api_version", "spec_replicas", "spec_strategy_type",
		"spec_paused", "spec_revision_history_limit", "spec_progress_deadline_seconds", "spec_min_ready_seconds", "spec_rollbackto_revision",
		"spec_strategy_rolling_update_max_surge", "spec_strategy_rolling_update_max_unavailable"})

	MachineDeploymentStatusCondition = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_deployment_status_condition",
		Help: "Information of the mcm managed Machinedeployments' status conditions.",
	}, []string{"name", "namespace", "uid", "condition", "status"})

	MachineDeploymentStatusAvailableReplicas = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_deployment_status_available_replicas",
		Help: "Count of the mcm managed Machinedeployments available replicas.",
	}, []string{"name", "namespace", "uid"})
	MachineDeploymentStatusUnavailableReplicas = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_deployment_status_unavailable_replicas",
		Help: "Count of the mcm managed Machinedeployments unavailable replicas.",
	}, []string{"name", "namespace", "uid"})
	MachineDeploymentStatusReadyReplicas = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_deployment_status_ready_replicas",
		Help: "Count of the mcm managed Machinedeployments ready replicas.",
	}, []string{"name", "namespace", "uid"})
	MachineDeploymentStatusUpdatedReplicas = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_deployment_status_updated_replicas",
		Help: "Count of the mcm managed Machinedeployments updated replicas.",
	}, []string{"name", "namespace", "uid"})
	MachineDeploymentStatusCollisionCount = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_deployment_status_collision_count",
		Help: "Mcm managed Machinedeployments collision count.",
	}, []string{"name", "namespace", "uid"})
	MachineDeploymentStatusReplicas = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_deployment_status_replicas",
		Help: "Count of the mcm managed Machinedeployments replicas.",
	}, []string{"name", "namespace", "uid"})

	MachineDeploymentStatusFailedMachines = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name: "mcm_machine_deployment_failed_machines",
		Help: "Information of the mcm managed Machinedeployments' failed machines.",
	}, []string{"name", "namespace", "uid", "failed_machine_name", "failed_machine_provider_id", "failed_machine_owner_ref",
		"failed_machine_last_operation_state",
		"failed_machine_last_operation_machine_operation_type"})

	ApiRequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mcm_cloud_api_requests_total",
		Help: "Number of Cloud Service API requests, partitioned by provider, and service.",
	}, []string{"provider", "service"},
	)
	ApiFailedRequestCount = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mcm_cloud_api_requests_failed_total",
		Help: "Number of Failed Cloud Service API requests, partitioned by provider, and service.",
	}, []string{"provider", "service"},
	)

	// ScrapeFailedCounter is a Prometheus metric, which counts errors during metrics collection.
	ScrapeFailedCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "mcm_scrape_failure_total",
		Help: "Total count of scrape failures.",
	}, []string{"kind"})
)

func init() {
	prometheus.MustRegister(ScrapeFailedCounter)
	prometheus.MustRegister(MachineCreated)
	prometheus.MustRegister(MachineInfo)
	prometheus.MustRegister(MachineStatusCondition)
	prometheus.MustRegister(MachineCSPhase)
	prometheus.MustRegister(MachineSetCreated)
	prometheus.MustRegister(MachineSetInfo)
	prometheus.MustRegister(MachineSetInfoSpecReplicas)
	prometheus.MustRegister(MachineSetInfoSpecMinReadySeconds)
	prometheus.MustRegister(MachineSetStatus)
	prometheus.MustRegister(MachineSetStatusCondition)
	prometheus.MustRegister(MachineSetStatusFailedMachines)
	prometheus.MustRegister(MachineDeploymentCreated)
	prometheus.MustRegister(MachineDeploymentInfo)
	prometheus.MustRegister(MachineDeploymentStatusCondition)
	prometheus.MustRegister(MachineDeploymentStatusAvailableReplicas)
	prometheus.MustRegister(MachineDeploymentStatusUnavailableReplicas)
	prometheus.MustRegister(MachineDeploymentStatusReadyReplicas)
	prometheus.MustRegister(MachineDeploymentStatusUpdatedReplicas)
	prometheus.MustRegister(MachineDeploymentStatusCollisionCount)
	prometheus.MustRegister(MachineDeploymentStatusReplicas)
	prometheus.MustRegister(MachineDeploymentStatusFailedMachines)
	prometheus.MustRegister(ApiRequestCount)
	prometheus.MustRegister(ApiFailedRequestCount)
}
package utils

const (
	ResourceName = "ucloud.cn/gpu-mem"
	CountName    = "ucloud.cn/gpu-count"

	EnvNVGPU              = "NVIDIA_VISIBLE_DEVICES"
	EnvResourceIndex      = "UCLOUD_CN_GPU_MEM_IDX" //满足GPU资源条件且剩余最少的GPU卡，保存到Pod的annotation中
	EnvResourceByPod      = "UCLOUD_CN_GPU_MEM_POD" //该Pod申请的GPU Memory，保存到Pod的annotation中
	EnvResourceByDev      = "UCLOUD_CN_GPU_MEM_DEV" //每块卡的内存数
	EnvAssignedFlag       = "UCLOUD_CN_GPU_MEM_ASSIGNED" //它被初始化为false。它表示该Pod在调度时刻被指定到了某块GPU卡，但是并没有真正在节点上创建该Pod
	EnvResourceAssumeTime = "UCLOUD_CN_GPU_MEM_ASSUME_TIME" //申请GPU Memory的时间，保存到Pod的annotation中
)
//  发现分配节点上没有GPU资源符合条件，此时不进行绑定，直接不报错退出，默认调度器会在assume超时后重新调度
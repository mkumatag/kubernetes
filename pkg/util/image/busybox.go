package image

var busybox = MultiArchImage{
	AMD64:   "busybox",
	ARM:     "armhf/busybox",
	ARM64:   "aarch64/busybox",
	PPC64LE: "ppc64le/busybox",
}

func getBusyBoxImage() string {
	return busybox.getArchImage()
}

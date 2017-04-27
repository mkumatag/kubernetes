package image

var ubuntu = MultiArchImage{
	AMD64:"ubuntu",
	ARM:     "armhf/ubuntu",
	ARM64:   "aarch64/ubuntu",
	PPC64LE: "ppc64le/ubuntu",
}

func getUbuntuImage() string {
	return ubuntu.getArchImage()
}

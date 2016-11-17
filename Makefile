export UDIR= .
export GOC = x86_64-xen-ethos-6g
export GOL = x86_64-xen-ethos-6l
export ETN2GO = etn2go
export ET2G   = et2g
export EG2GO  = eg2go

export GOARCH = amd64
export TARGET_ARCH = x86_64
export GOETHOSINCLUDE=/usr/lib64/go/pkg/ethos_$(GOARCH)
export GOLINUXINCLUDE=/usr/lib64/go/pkg/linux_$(GOARCH)


install.rootfs = /var/lib/ethos/ethos-default-$(TARGET_ARCH)/rootfs
install.minimaltd.rootfs = /var/lib/ethos/minimaltd/rootfs


.PHONY: all install
all: Tree

install: Tree
	ethosTypeInstall $(install.rootfs) $(install.minimaltd.rootfs) FileType
	install Tree $(install.rootfs)/programs
	install Tree $(install.minimaltd.rootfs)/programs
	echo -n /programs/Tree | ethosStringEncode > $(install.rootfs)/etc/init/console
	mkdir -p $(install.rootfs)/root/EthosProject
	setfattr -n user.ethos.typeHash -v $(shell egPrint FileType hash FileType) $(install.rootfs)/root/EthosProject


FileType.go: FileType.t
	$(ETN2GO) . FileType main $^

Tree: Tree.go FileType.go
	ethosGo $^ 

clean:
	rm -rf FileTypeIndex
	rm -rf FileType.go
	rm -rf Tree
	rm -rf Tree.goo.ethos
	rm -rf FileType

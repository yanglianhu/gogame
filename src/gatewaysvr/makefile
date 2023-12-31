
APP := LittleServerAlian
TARGET := gatewaysvr
CONFIG := ../../run/binconf
tarstoken := eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1aWQiOiJhZG1pbiIsImlhdCI6MTY2OTI1Njc3MSwiZXhwIjoxNzQwNzk1OTcxfQ.wuajor2WoA7yrMGR61T48NUVCC7BTa1Y8IAcTYcgU2M
tarsurl:= http://10.12.6.51:3000
tarsapp:=LittleServerAlian
GO_BUILD_FLAG:=-gcflags "-N -l"
export LC_ALL   = en_US.UTF-8
export LANG     = en_US.UTF-8
#export GOPATH = $(shell pwd):$(shell pwd)/../../gopath

#-------------------------------------------------------------------------------

GO          = go

#-------------------------------------------------------------------------------
TARS2GO     := tars2go
J2GO_FLAG   := -outdir=scripts ${J2GO_FLAG}
#-------------------------------------------------------------------------------

TARS_SRC     := $(wildcard *.tars)
GO_SRC      := $(wildcard ./scripts/*.go)

#----------------------------------------------------------------------------------


copyfile = if test -z "$(APP)" || test -z "$(TARGET)"; then \
               echo "['APP' or 'TARGET' option is empty.]"; exit 1; \
           	else \
		       	if test ! -d $(2); then \
              		echo "[No such dir:$(2), now we create it.]";\
    				mkdir -p $(2);\
				fi; \
         		echo "[Copy file $(1) -> $(2)]"; \
         		cp -v $(1) $(2); \
			fi;

ALL: $(TARGET)
#----------------------------------------------------------------------------------
$(TARGET): TARSBUILD $(GO_SRC)
	$(GO) build $(GO_BUILD_FLAG) -o $@ scripts/main.go
#----------------------------------------------------------------------------------
TARSBUILD: $(TARS_SRC)
	$(TARS2GO) $(J2GO_FLAG) $(TARS_SRC)
#----------------------------------------------------------------------------------
tar: $(TARGET) $(CONFIG)
	@if [ -d $(TARGET)_tmp_dir ]; then \
		echo "dir has exist:$(TARGET)_tmp_dir, abort."; \
		exit 1; \
	else \
		mkdir $(TARGET)_tmp_dir $(TARGET)_tmp_dir/$(TARGET);\
		cp -rf $(TARGET) $(CONFIG) $(TARGET)_tmp_dir/$(TARGET)/; \
		cd $(TARGET)_tmp_dir; tar --exclude=".svn" --exclude="_svn" -czvf $(TARGET).tgz $(TARGET)/; cd ..; \
		if [ -f "$(TARGET).tgz" ]; then \
			mv -vf $(TARGET).tgz $(TARGET).`date +%Y%m%d%H%M%S`.tgz; \
		fi; \
		mv $(TARGET)_tmp_dir/$(TARGET).tgz ./; \
		rm -rf $(TARGET)_tmp_dir; \
		echo "tar cvfz $(TARGET).tgz ..."; \
	fi
pub:$(tar) $(TARGET) $(CONFIG)
	make tar
	curl -X POST ${tarsurl}/api/upload_and_publish?ticket=$(tarstoken) -Fsuse=@$(TARGET).tgz -Fapplication=$(tarsapp) -Fmodule_name=gatewaysvr -Fcomment=jenkins
HELP += $(HELP_TAR)

ifneq ($(TARS_SRC),)

SERVER_NAME := $(TARGET)

endif
#----------------------------------------------------------------------------------

clean:
	rm -vf $(DEPEND_TARS_OBJ) $(INVOKE_DEPEND_TARS_OBJ) $(LOCAL_OBJ) $(TARGET) $(TARGETS) $(DEP_FILE) ${CLEANFILE} .*.d.tmp gmon.out
	rm -vf *$(TARGET)*.tgz

cleanall:
	rm -vf $(DEPEND_TARS_H) $(DEPEND_TARS_CPP) $(DEPEND_TARS_OBJ) $(LOCAL_OBJ) $(HCE_H) $(HCE_CPP) $(TARGET) $(TARGETS) $(DEP_FILE) ${CLEANFILE} *.o .*.d.tmp .*.d gmon.out
	rm -vf *$(TARGET)*.tgz
stop:
	cp -rf $(TARGET) /usr/local/app/tars/tarsnode/data/$(APP).$(TARGET)/bin/.
	cp -rf $(CONFIG) /usr/local/app/tars/tarsnode/data/$(APP).$(TARGET)/bin/.
	ps -efww|grep -w /usr/local/app/tars/tarsnode/data/$(APP).$(TARGET)/bin|grep -v grep|cut -c 9-15|xargs kill -9
pub2web:
	python3 ../../tool/upload_serverpkg.py $(TARGET) svntest $(TARGET).tgz

HELP += $(HELP_CLEAN)
HELP += $(HELP_CLEANALL)

HELP_CLEAN    = "\n\e[1;33mclean\e[0m:\t\t[remove $(LOCAL_OBJ) $(TARGET)]"
HELP_CLEANALL = "\n\e[1;33mcleanall\e[0m:\t[clean & rm .*.d]"
HELP_TAR      = "\n\e[1;33mtar\e[0m:\t\t[will do 'tar $(TARGET).tgz $(RELEASE_FILE)']"

help:
	@echo -e $(HELP)"\n"

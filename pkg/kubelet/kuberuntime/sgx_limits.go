package kuberuntime

//#include <fcntl.h>
//#include <stdio.h>
//#include <unistd.h>
//#include <sys/ioctl.h>
//#include <string.h>
//
//struct sgx_enclave_limit {
//    char cgroup_path[200];
//    unsigned int max_pages;
//} __packed;
//
//int communicate_limits(char* cgroup_name, unsigned int limit) {
//    int fd = open("/dev/isgx", O_RDWR);
//    if (fd < 0) {
//        printf("Error: open isgx device failed\n");
//        return -1;
//    }
//    struct sgx_enclave_limit limits_struct;
//    strncpy(limits_struct.cgroup_path, cgroup_name, 200);
//    limits_struct.max_pages = limit;
//
//    return ioctl(fd, 1087153156, &limits_struct);
//}
import "C"

func CommunicateLimitsToSgxDriver(cgroupName string, limit uint32) {
	C.communicate_limits(C.CString(cgroupName), C.uint(limit))
}

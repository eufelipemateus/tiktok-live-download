package process

/*
#cgo CXXFLAGS: --std=c++11
#cgo CXXFLAGS: -I/usr/include/opencv4
#cgo LDFLAGS: -lopencv_core -lopencv_imgproc -lopencv_highgui -lopencv_videoio
#cgo LDFLAGS: ./lib/libvideo_processor.so -ldl
#include <stdlib.h>

extern int processVideo(const char* videoInput, const char* videoOutput, const char* viewFile, const char* chatFile);
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func GenerateVideoFinal(inputFile string, outputFile string, viewFile string, chatFile string) {

	videoInputC := C.CString(inputFile)
	videoOutputC := C.CString(outputFile)
	viewFileC := C.CString(viewFile)
	chatFileC := C.CString(chatFile)


	defer C.free(unsafe.Pointer(videoInputC))
	defer C.free(unsafe.Pointer(videoOutputC))
	defer C.free(unsafe.Pointer(viewFileC))
	defer C.free(unsafe.Pointer(chatFileC))

	result := C.processVideo(videoInputC, videoOutputC, viewFileC,chatFileC)
	if result != 0 {
		fmt.Println("Erro ao processar o vídeo")
	} else {
		fmt.Println("Vídeo processado com sucesso!")
	}
}


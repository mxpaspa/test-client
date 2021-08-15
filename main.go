package main

import ( // "encoding/json"
	// "fmt"
	// "io/ioutil"
	// "log"
	// "net/http"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	// "github.com/redhat-marketplace/redhat-marketplace-operator/v2/apis/marketplace/v1beta1"
	olmv1alpha1 "github.com/operator-framework/api/pkg/operators/v1alpha1"
	"github.com/redhat-marketplace/redhat-marketplace-operator/v2/apis/marketplace/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// car csv

// func main() {

//     resp, err := http.Get("http://localhost:8100/get-templates")
//     if err != nil {
//         log.Fatalf("error on get %s",err)
//     }

//     if resp.StatusCode == http.StatusNoContent {
//         fmt.Println(resp.Status)
//         return
//     }

//     data, err := ioutil.ReadAll(resp.Body)
// 	defer resp.Body.Close()

// 	if err != nil {
// 		log.Fatalf("error reading body %s",err)
// 	}

//     fmt.Println("status code",resp.StatusCode)
//     fmt.Println("status",resp.Status)

//     templates := []v1beta1.MeterDefinition{}

//     err = json.Unmarshal(data,&templates)
//     if err != nil {
// 		log.Fatalf("error unmarshalling %s",err)
//     }

//     fUtils.PrettyPrint(templates)

// }

func main(){
    csv := &olmv1alpha1.ClusterServiceVersion{
        ObjectMeta: metav1.ObjectMeta{
            Name:      "testName",
            Namespace: "testNamespace",
        },
    }
    	// marshal CSV struct o JSON
	requestBody, err := json.Marshal(csv)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8100/get-system-meterdefs/test",
		"application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatal(err)
	}

    data, err := ioutil.ReadAll(resp.Body)
    defer resp.Body.Close()

    if err != nil {
        log.Fatalf("error reading body %s",err)
    }

    fmt.Println("status code",resp.StatusCode)
    fmt.Println("status",resp.Status)

    templates := []v1beta1.MeterDefinition{}

    err = json.Unmarshal(data,&templates)
    if err != nil {
        log.Fatalf("error unmarshalling %s",err)
    }
    
    PrettyPrint(templates)
    // fUtils.PrettyPrint(templates)
}

func PrettyPrint(in interface{}) {
	out, _ := json.MarshalIndent(in, "", "    ")
	println(string(out))
}

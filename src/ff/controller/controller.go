package controller

import (
    "reflect"
    "errors"
    "fmt"
    "net/http"
)

var DefaultController = &defaultController
var defaultController controller

func init() {
    DefaultController.controllerPool = make(map[string]reflect.Type)
}

type controller struct {
  controllerPool map[string]reflect.Type
}

func Register(controllerName string, controllerInstance ControllerInterface) {
    DefaultController.controllerPool[controllerName] = reflect.TypeOf(controllerInstance)
}

func GetController(controllerName string) (reflect.Type, error) {
    controllerType, exists := DefaultController.controllerPool[controllerName]
    if !exists {
        return nil, errors.New("controller no exists")
    }

    return controllerType, nil
}

func Do(controllerName string, actionName string, resp http.ResponseWriter, r *http.Request) error {
    controllerType, err := GetController(controllerName)
    if err != nil {
        return err
    }

    tmp := []byte(actionName + "Action")
    tmp[0] -= 32
    actionName = string(tmp)
    
    ctrlVar := reflect.New(controllerType)
    actMethod := ctrlVar.MethodByName(actionName)
    if !actMethod.IsValid() {
        return errors.New("action not found")
    }
    in := []reflect.Value{reflect.ValueOf(resp), reflect.ValueOf(r)}
    actMethod.Call(in)
    
    fmt.Println("contorller name:", controllerName, "controller Type:", controllerType, ", action name:", actionName)
    return nil
}

type ControllerInterface interface {}

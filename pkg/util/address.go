package util

import (
	"errors"

	listenersv1alpha1 "github.com/zncdata-labs/listener-operator/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
)

type AddressInfo struct {
	Address     string
	AddressType listenersv1alpha1.AddressType
}

type ListenerIngress struct {
	AddressInfo
	Ports []listenersv1alpha1.PortSpec
}

func GetPriorNodeAddress(node *corev1.Node) (*AddressInfo, error) {
	for _, address := range node.Status.Addresses {
		if address.Type == corev1.NodeExternalIP {
			return &AddressInfo{
				Address:     address.Address,
				AddressType: listenersv1alpha1.AddressTypeIP,
			}, nil
		} else if address.Type == corev1.NodeInternalIP {
			return &AddressInfo{
				Address:     address.Address,
				AddressType: listenersv1alpha1.AddressTypeIP,
			}, nil
		} else if address.Type == corev1.NodeHostName {
			return &AddressInfo{
				Address:     address.Address,
				AddressType: listenersv1alpha1.AddressTypeHostname,
			}, nil
		}
	}
	return nil, errors.New("no address found")
}

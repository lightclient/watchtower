package kubernetes

import (
	"fmt"
	"strings"

	root "github.com/c-o-l-o-r/watchtower/manager/pkg"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WatchtowerService struct {
	client *Client
}

func NewWatchtowerService(c *Client) *WatchtowerService {
	return &WatchtowerService{c}
}

func (p *WatchtowerService) CreateWatchtower(a root.WatchtowerAttributes) error {

	// Address provided to kubernetes must be lowercase in order to adhere to RFC 1123
	lowercaseAddress := strings.ToLower(a.Address)

	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("watchtower-%s", lowercaseAddress),
			Namespace: "watchtower",
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "watchtower",
				},
			},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "watchtower",
					},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Name:    "web",
							Image:   "mattgarnett/watchtower:600ba50",
							Command: []string{"node"},
							Args:    []string{"./bin/cli.js", "-p", a.Phone},
							Ports: []apiv1.ContainerPort{
								{
									Name:          "http",
									Protocol:      apiv1.ProtocolTCP,
									ContainerPort: 80,
								},
							},
							Env: []apiv1.EnvVar{
								{
									Name: "TWILIO_AUTH_TOKEN",
									ValueFrom: &apiv1.EnvVarSource{
										ConfigMapKeyRef: &apiv1.ConfigMapKeySelector{
											LocalObjectReference: apiv1.LocalObjectReference{
												Name: "watchtower-config",
											},
											Key: "TWILIO_AUTH_TOKEN",
										},
									},
								},
								{
									Name: "TWILIO_SID",
									ValueFrom: &apiv1.EnvVarSource{
										ConfigMapKeyRef: &apiv1.ConfigMapKeySelector{
											LocalObjectReference: apiv1.LocalObjectReference{
												Name: "watchtower-config",
											},
											Key: "TWILIO_SID",
										},
									},
								},
								{
									Name: "TWILIO_FROM_NUMBER",
									ValueFrom: &apiv1.EnvVarSource{
										ConfigMapKeyRef: &apiv1.ConfigMapKeySelector{
											LocalObjectReference: apiv1.LocalObjectReference{
												Name: "watchtower-config",
											},
											Key: "TWILIO_FROM_NUMBER",
										},
									},
								},
								{
									Name: "TWITTER_CONSUMER_KEY",
									ValueFrom: &apiv1.EnvVarSource{
										ConfigMapKeyRef: &apiv1.ConfigMapKeySelector{
											LocalObjectReference: apiv1.LocalObjectReference{
												Name: "watchtower-config",
											},
											Key: "TWITTER_CONSUMER_KEY",
										},
									},
								},
								{
									Name: "TWITTER_CONSUMER_SECRET",
									ValueFrom: &apiv1.EnvVarSource{
										ConfigMapKeyRef: &apiv1.ConfigMapKeySelector{
											LocalObjectReference: apiv1.LocalObjectReference{
												Name: "watchtower-config",
											},
											Key: "TWITTER_CONSUMER_SECRET",
										},
									},
								},
								{
									Name: "TWITTER_ACCESS_TOKEN",
									ValueFrom: &apiv1.EnvVarSource{
										ConfigMapKeyRef: &apiv1.ConfigMapKeySelector{
											LocalObjectReference: apiv1.LocalObjectReference{
												Name: "watchtower-config",
											},
											Key: "TWITTER_ACCESS_TOKEN",
										},
									},
								},
								{
									Name: "TWITTER_ACCESS_TOKEN_SECRET",
									ValueFrom: &apiv1.EnvVarSource{
										ConfigMapKeyRef: &apiv1.ConfigMapKeySelector{
											LocalObjectReference: apiv1.LocalObjectReference{
												Name: "watchtower-config",
											},
											Key: "TWITTER_ACCESS_TOKEN_SECRET",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	_, err := p.client.deployments.Create(deployment)
	return err
}

func int32Ptr(i int32) *int32 { return &i }

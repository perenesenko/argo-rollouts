package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestHashUtils(t *testing.T) {
	templateRed := generatePodTemplate("red")

	hashRed := ComputePodTemplateHash(&templateRed, nil)

	t.Run("HashForSameTemplates", func(t *testing.T) {
		template := generatePodTemplate("red")
		podHash := ComputePodTemplateHash(&template, nil)
		assert.Equal(t, hashRed, podHash)
	})
	t.Run("HashForDifferentTemplates", func(t *testing.T) {
		template := generatePodTemplate("blue")
		podHash := ComputePodTemplateHash(&template, nil)
		assert.NotEqual(t, hashRed, podHash)
	})
}

func generatePodTemplate(image string) corev1.PodTemplateSpec {
	podLabels := map[string]string{"name": image}

	return corev1.PodTemplateSpec{
		ObjectMeta: metav1.ObjectMeta{
			Labels: podLabels,
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:                   image,
					Image:                  image,
					ImagePullPolicy:        corev1.PullAlways,
					TerminationMessagePath: corev1.TerminationMessagePathDefault,
				},
			},
			DNSPolicy:       corev1.DNSClusterFirst,
			RestartPolicy:   corev1.RestartPolicyAlways,
			SecurityContext: &corev1.PodSecurityContext{},
		},
	}
}

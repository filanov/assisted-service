package eventwrapper

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/jinzhu/gorm"
	"github.com/openshift/assisted-service/internal/common"

	"github.com/go-openapi/strfmt"
	"github.com/openshift/assisted-service/internal/events"
	"sigs.k8s.io/controller-runtime/pkg/event"
)

type Wrapper struct {
	NotificationsChan chan event.GenericEvent
	Inner             events.Handler
	DB                *gorm.DB
	Log               logrus.FieldLogger
}

func (w *Wrapper) AddEvent(ctx context.Context, clusterID strfmt.UUID, hostID *strfmt.UUID, severity string, msg string, eventTime time.Time) {
	w.Inner.AddEvent(ctx, clusterID, hostID, severity, msg, eventTime)

	if hostID == nil {
		c := &common.Cluster{}
		if err := w.DB.Take(c, "id = ?", clusterID).Error; err != nil {
			return
		}
		w.NotificationsChan <- event.GenericEvent{
			Meta: &metav1.ObjectMeta{
				Name:      c.KubeKeyName,
				Namespace: c.KubeKeyNamespace,
			},
		}
	}
}

func (w *Wrapper) GetEvents(clusterID strfmt.UUID, hostID *strfmt.UUID) ([]*events.Event, error) {
	return w.Inner.GetEvents(clusterID, hostID)
}

func (w *Wrapper) DeleteClusterEvents(clusterID strfmt.UUID) {
	w.Inner.DeleteClusterEvents(clusterID)
}

package public

import (
	timesPb ""

	"context"
	"io"
)

type TimesheetRepository interface {
	CreateTimesheet(ctx context.Context, rq *timesPb.CreateTimesheetRequest) (*timesPb.Timesheet, error)
	UpdateTimesheet(ctx context.Context, rq *timesPb.UpdateTimesheetRequest) (*timesPb.Timesheet, error)
	GetTimesheet(ctx context.Context, rq *timesPb.TimesheetIdRequest) (*timesPb.Timesheet, error)
	SearchTimesheet(ctx context.Context, rq *timesPb.SearchTimesheetRequest) (*timesPb.SearchTimesheetResponse, error)
	DeleteTimesheet(ctx context.Context, rq *timesPb.TimesheetIdRequest) error
	CreateEvent(ctx context.Context, rq *timesPb.CreateEventRequest) (*timesPb.Event, error)
	UpdateEvent(ctx context.Context, rq *timesPb.UpdateEventRequest) (*timesPb.Event, error)
	GetEvent(ctx context.Context, rq *timesPb.EventIdRequest) (*timesPb.Event, error)
	DeleteEvent(ctx context.Context, rq *timesPb.EventIdRequest) error
	SearchEvents(ctx context.Context, rq *timesPb.TimesheetIdRequest) (*timesPb.SearchResponse, error)
}

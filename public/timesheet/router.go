package timesheet

import ()

type Router struct {
	c            Controller
	routeBuilder *routing.RouteBuilder
}

func (r *Router) Set() error {
	return r.routeBuilder.Build(
		routing.R("/api/timesheet/timetable", r.c.CreateTimesheet).POST().Authorize(),
		routing.R("/api/timesheet/timetable", r.c.SearchTimesheet).GET().Authorize(),
		routing.R("/api/timesheet/timetable/{id}", r.c.UpdateTimesheet).PUT().Authorize(),
		routing.R("/api/timesheet/timetable/{id}", r.c.GetTimesheet).GET().Authorize(),
		routing.R("/api/timesheet/timetable/{id}", r.c.DeleteTimesheet).DELETE().Authorize(),
		routing.R("/api/timesheet/event", r.c.CreateEvent).POST().Authorize(),
		routing.R("/api/timesheet/event", r.c.SearchEvents).GET().Authorize(),
		routing.R("/api/timesheet/event/{id}", r.c.UpdateEvent).PUT().Authorize(),
		routing.R("/api/timesheet/event/{id}", r.c.GetEvent).GET().Authorize(),
		routing.R("/api/timesheet/event/{id}", r.c.DeleteEvent).DELETE().Authorize(),
	)
}

func NewRouter(c Controller, routeBuilder *routing.RouteBuilder) kitHttp.RouteSetter {
	return &Router{
		c:            c,
		routeBuilder: routeBuilder,
	}
}

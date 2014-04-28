Recurrence
==========

Calendar recurrence rules in Go.

Basically, implementing the strategy described in [Martin Fowler's paper](http://martinfowler.com/apsupp/recurring.pdf). Read it. It's fun.

## Schedule

The `Schedule` interface is the foundation of the recurrence package. By using and combining schedules, we can represent all kinds of recurrence rules.

`schedule.IsOccurring(time.Time) bool` - does this `time.Time` occur in this schedule?
`schedule.Occurrences(TimeRange) chan time.Time` - generate `time.Time`s occuring with the passed `TimeRange`

## Day

Integer day of the month, 1 through 31, or the constant `Last`.

```go
first := recurrence.Day(First)
last := recurrence.Day(Last)
twentieth := recurrence.Day(20)
```

## Week

Integer week of the month, 1 through 5, or the constant `Last`.

```go
first := recurrence.Week(First)
last := recurrence.Week(Last)
third := recurrence.Week(Third)
```

## Weekday

Day of the week, Sunday through Monday. Constants are defined so you can use them with ease.

```go
recurrence.Sunday.IsOccurring(time.Now())
```

## Month

A month of the year. Constants are defined to be used with ease.

```go
recurrence.January.IsOccurring(time.Now())
```

## Set Operations

### Intersection

Intersection is a slice of Schedules. `IsOccurring` is only satisfied if all members of the slice are true. (Set intersection).

```go
// Complex Rules
american_thanksgiving := recurrence.Intersection{recurrence.Week(4), recurrence.Thursday, recurrence.November}
```

### Union

Union is a slice of Schedules. `IsOccurring` is satisfied if any member of the slice is occurring. (Set union).

```go
weekends := recurrence.Union{recurrence.Saturday, recurrence.Sunday}
```

### Exclusion

Exclusion computes the set difference between two schedules.

```go
every_friday_except_the_last := recurrence.Exclusion{
  Schedule: recurrence.Friday,
  Exclude: recurrence.Week(recurrence.Last)
}
```


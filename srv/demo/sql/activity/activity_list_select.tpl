select
    example_activity.id
from
	example_activity
where
	example_activity.id {{.ne}} ""
LIMIT {{.size}} OFFSET {{.offset}}
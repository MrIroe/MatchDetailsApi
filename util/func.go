package util

import (
	obj "matchDetailsApi/objects"
)

func AverageDeltas(timeline obj.ParticipantTimelineDto) (float64, float64, float64) {
	creepsPerMin := (timeline.CreepsPerMinDeltas["0-10"] + timeline.CreepsPerMinDeltas["10-20"] + timeline.CreepsPerMinDeltas["20-30"]) / 3
	csDiff := (timeline.CsDiffPerMinDeltas["0-10"] + timeline.CsDiffPerMinDeltas["10-20"] + timeline.CsDiffPerMinDeltas["20-30"]) / 3
	expDiff := (timeline.XpDiffPerMinDeltas["0-10"] + timeline.XpDiffPerMinDeltas["10-20"] + timeline.XpDiffPerMinDeltas["20-30"]) / 3

	return creepsPerMin, csDiff, expDiff
}

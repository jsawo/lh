package takeout

import "time"

type LocationDataFile struct {
	TimelineObjects []TimelineObject `json:"timelineObjects"`
}

type TimelineObject struct {
	PlaceVisit      PlaceVisit      `json:"placeVisit,omitempty"`
	ActivitySegment ActivitySegment `json:"activitySegment,omitempty"`
}

// generated with https://mholt.github.io/json-to-go/

// PlaceVisit
type PlaceVisit struct {
	Location                Location                  `json:"location"`
	Duration                Duration                  `json:"duration"`
	PlaceConfidence         string                    `json:"placeConfidence"`
	CenterLatE7             int                       `json:"centerLatE7"`
	CenterLngE7             int                       `json:"centerLngE7"`
	VisitConfidence         int                       `json:"visitConfidence"`
	OtherCandidateLocations []OtherCandidateLocations `json:"otherCandidateLocations"`
	EditConfirmationStatus  string                    `json:"editConfirmationStatus"`
	SimplifiedRawPath       SimplifiedRawPath         `json:"simplifiedRawPath"`
	LocationConfidence      int                       `json:"locationConfidence"`
	PlaceVisitType          string                    `json:"placeVisitType"`
	PlaceVisitImportance    string                    `json:"placeVisitImportance"`
}
type SourceInfo struct {
	DeviceTag int `json:"deviceTag"`
}
type Location struct {
	LatitudeE7            int        `json:"latitudeE7,omitempty"`
	LongitudeE7           int        `json:"longitudeE7,omitempty"`
	PlaceID               string     `json:"placeId,omitempty"`
	Address               string     `json:"address,omitempty"`
	Name                  string     `json:"name,omitempty"`
	SourceInfo            SourceInfo `json:"sourceInfo,omitempty"`
	LocationConfidence    float64    `json:"locationConfidence,omitempty"`
	CalibratedProbability float64    `json:"calibratedProbability,omitempty"`
	AccuracyMetres        int        `json:"accuracyMetres,omitempty"`
}
type Duration struct {
	StartTimestamp time.Time `json:"startTimestamp"`
	EndTimestamp   time.Time `json:"endTimestamp"`
}
type OtherCandidateLocations struct {
	LatitudeE7            int     `json:"latitudeE7"`
	LongitudeE7           int     `json:"longitudeE7"`
	PlaceID               string  `json:"placeId"`
	Address               string  `json:"address"`
	Name                  string  `json:"name,omitempty"`
	LocationConfidence    float64 `json:"locationConfidence"`
	CalibratedProbability float64 `json:"calibratedProbability"`
	SemanticType          string  `json:"semanticType,omitempty"`
}
type Point struct {
	LatE7          int       `json:"latE7"`
	LngE7          int       `json:"lngE7"`
	AccuracyMeters int       `json:"accuracyMeters"`
	Timestamp      time.Time `json:"timestamp"`
}
type SimplifiedRawPath struct {
	Points         []Point `json:"points"`
	Source         string  `json:"source"`
	DistanceMeters float64 `json:"distanceMeters"`
}

// ActivitySegment

type ActivitySegment struct {
	StartLocation Location     `json:"startLocation"`
	EndLocation   Location     `json:"endLocation"`
	Duration      Duration     `json:"duration"`
	Distance      int          `json:"distance"`
	ActivityType  string       `json:"activityType"`
	Confidence    string       `json:"confidence"`
	Activities    []Activities `json:"activities"`
	WaypointPath  WaypointPath `json:"waypointPath"`
	ParkingEvent  ParkingEvent `json:"parkingEvent"`
}
type Activities struct {
	ActivityType string  `json:"activityType"`
	Probability  float64 `json:"probability"`
}
type Waypoint struct {
	LatE7 int `json:"latE7"`
	LngE7 int `json:"lngE7"`
}
type RoadSegment struct {
	PlaceID  string `json:"placeId"`
	Duration string `json:"duration"`
}
type WaypointPath struct {
	Waypoints      []Waypoint    `json:"waypoints"`
	Source         string        `json:"source"`
	RoadSegment    []RoadSegment `json:"roadSegment"`
	DistanceMeters float64       `json:"distanceMeters"`
	TravelMode     string        `json:"travelMode"`
	Confidence     float64       `json:"confidence"`
}
type ParkingEvent struct {
	Location       Location  `json:"location"`
	Method         string    `json:"method"`
	LocationSource string    `json:"locationSource"`
	Timestamp      time.Time `json:"timestamp"`
}

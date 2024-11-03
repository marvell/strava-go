package strava

import (
	"time"
)

// Models generated from Strava API v3 Swagger specification

// ActivityStats represents the stats of an athlete's activities.
type ActivityStats struct {
	BiggestRideDistance       float64        `json:"biggest_ride_distance"`
	BiggestClimbElevationGain float64        `json:"biggest_climb_elevation_gain"`
	RecentRideTotals          *ActivityTotal `json:"recent_ride_totals"`
	RecentRunTotals           *ActivityTotal `json:"recent_run_totals"`
	RecentSwimTotals          *ActivityTotal `json:"recent_swim_totals"`
	YtdRideTotals             *ActivityTotal `json:"ytd_ride_totals"`
	YtdRunTotals              *ActivityTotal `json:"ytd_run_totals"`
	YtdSwimTotals             *ActivityTotal `json:"ytd_swim_totals"`
	AllRideTotals             *ActivityTotal `json:"all_ride_totals"`
	AllRunTotals              *ActivityTotal `json:"all_run_totals"`
	AllSwimTotals             *ActivityTotal `json:"all_swim_totals"`
}

// ActivityTotal represents the total stats for an activity type.
type ActivityTotal struct {
	Count         int     `json:"count"`
	Distance      float64 `json:"distance"`
	MovingTime    int     `json:"moving_time"`
	ElapsedTime   int     `json:"elapsed_time"`
	ElevationGain float64 `json:"elevation_gain"`
}

// DetailedAthlete represents detailed information about an athlete.
type DetailedAthlete struct {
	ID                    uint          `json:"id"`
	Username              string        `json:"username"`
	ResourceState         int           `json:"resource_state"`
	Firstname             string        `json:"firstname"`
	Lastname              string        `json:"lastname"`
	City                  string        `json:"city"`
	State                 string        `json:"state"`
	Country               string        `json:"country"`
	Sex                   string        `json:"sex"`
	Premium               bool          `json:"premium"`
	Summit                bool          `json:"summit"`
	CreatedAt             time.Time     `json:"created_at"`
	UpdatedAt             time.Time     `json:"updated_at"`
	BadgeTypeID           uint          `json:"badge_type_id"`
	ProfileMedium         string        `json:"profile_medium"`
	Profile               string        `json:"profile"`
	Friend                interface{}   `json:"friend"`
	Follower              interface{}   `json:"follower"`
	FollowerCount         int           `json:"follower_count"`
	FriendCount           int           `json:"friend_count"`
	MutualFriendCount     int           `json:"mutual_friend_count"`
	AthleteType           int           `json:"athlete_type"`
	DatePreference        string        `json:"date_preference"`
	MeasurementPreference string        `json:"measurement_preference"`
	Clubs                 []SummaryClub `json:"clubs"`
	FTP                   int           `json:"ftp"`
	Weight                float64       `json:"weight"`
	Bikes                 []SummaryGear `json:"bikes"`
	Shoes                 []SummaryGear `json:"shoes"`
}

// SummaryClub represents summary information about a club.
type SummaryClub struct {
	ID              uint   `json:"id"`
	ResourceState   int    `json:"resource_state"`
	Name            string `json:"name"`
	ProfileMedium   string `json:"profile_medium"`
	Profile         string `json:"profile"`
	CoverPhoto      string `json:"cover_photo"`
	CoverPhotoSmall string `json:"cover_photo_small"`
	SportType       string `json:"sport_type"`
	City            string `json:"city"`
	State           string `json:"state"`
	Country         string `json:"country"`
	Private         bool   `json:"private"`
	MemberCount     int    `json:"member_count"`
	Featured        bool   `json:"featured"`
	Verified        bool   `json:"verified"`
	Url             string `json:"url"`
}

// SummaryGear represents summary information about gear (bike or shoes).
type SummaryGear struct {
	ID            string  `json:"id"`
	Primary       bool    `json:"primary"`
	Name          string  `json:"name"`
	ResourceState int     `json:"resource_state"`
	Distance      float64 `json:"distance"`
}

// Zones represents heart rate and power zones.
type Zones struct {
	HeartRate []ZoneRange `json:"heart_rate"`
	Power     []ZoneRange `json:"power"`
}

// ZoneRange represents a single zone range.
type ZoneRange struct {
	Min int `json:"min"`
	Max int `json:"max"`
}

// DetailedSegment represents detailed information about a segment.
type DetailedSegment struct {
	ID                  uint                `json:"id"`
	Name                string              `json:"name"`
	ActivityType        string              `json:"activity_type"`
	Distance            float64             `json:"distance"`
	AverageGrade        float64             `json:"average_grade"`
	MaximumGrade        float64             `json:"maximum_grade"`
	ElevationHigh       float64             `json:"elevation_high"`
	ElevationLow        float64             `json:"elevation_low"`
	StartLatlng         []float64           `json:"start_latlng"`
	EndLatlng           []float64           `json:"end_latlng"`
	ClimbCategory       int                 `json:"climb_category"`
	City                string              `json:"city"`
	State               string              `json:"state"`
	Country             string              `json:"country"`
	Private             bool                `json:"private"`
	Hazardous           bool                `json:"hazardous"`
	Starred             bool                `json:"starred"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
	TotalElevationGain  float64             `json:"total_elevation_gain"`
	Map                 PolylineMap         `json:"map"`
	EffortCount         int                 `json:"effort_count"`
	AthleteCount        int                 `json:"athlete_count"`
	StarCount           int                 `json:"star_count"`
	AthleteSegmentStats AthleteSegmentStats `json:"athlete_segment_stats"`
}

// PolylineMap represents a map with polyline data.
type PolylineMap struct {
	ID            string `json:"id"`
	Polyline      string `json:"polyline"`
	ResourceState int    `json:"resource_state"`
}

// AthleteSegmentStats represents an athlete's stats for a segment.
type AthleteSegmentStats struct {
	PRElapsedTime int       `json:"pr_elapsed_time"`
	PRDate        time.Time `json:"pr_date"`
	EffortCount   int       `json:"effort_count"`
}

// DetailedSegmentEffort represents detailed information about a segment effort.
type DetailedSegmentEffort struct {
	ID               uint            `json:"id"`
	ActivityID       uint            `json:"activity_id"`
	ElapsedTime      int             `json:"elapsed_time"`
	StartDate        time.Time       `json:"start_date"`
	StartDateLocal   time.Time       `json:"start_date_local"`
	Distance         float64         `json:"distance"`
	IsKOM            bool            `json:"is_kom"`
	Name             string          `json:"name"`
	Activity         SummaryActivity `json:"activity"`
	Athlete          SummaryAthlete  `json:"athlete"`
	MovingTime       int             `json:"moving_time"`
	StartIndex       int             `json:"start_index"`
	EndIndex         int             `json:"end_index"`
	AverageCadence   float64         `json:"average_cadence"`
	AverageWatts     float64         `json:"average_watts"`
	DeviceWatts      bool            `json:"device_watts"`
	AverageHeartrate float64         `json:"average_heartrate"`
	MaxHeartrate     float64         `json:"max_heartrate"`
	Segment          SummarySegment  `json:"segment"`
	KOMRank          int             `json:"kom_rank"`
	PRRank           int             `json:"pr_rank"`
	Hidden           bool            `json:"hidden"`
}

// SummaryActivity represents summary information about an activity.
type SummaryActivity struct {
	ID            uint `json:"id"`
	ResourceState int  `json:"resource_state"`
}

// SummaryAthlete represents summary information about an athlete.
type SummaryAthlete struct {
	ID            uint `json:"id"`
	ResourceState int  `json:"resource_state"`
}

// SummarySegment represents summary information about a segment.
type SummarySegment struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	ActivityType  string    `json:"activity_type"`
	Distance      float64   `json:"distance"`
	AverageGrade  float64   `json:"average_grade"`
	MaximumGrade  float64   `json:"maximum_grade"`
	ElevationHigh float64   `json:"elevation_high"`
	ElevationLow  float64   `json:"elevation_low"`
	StartLatlng   []float64 `json:"start_latlng"`
	EndLatlng     []float64 `json:"end_latlng"`
	ClimbCategory int       `json:"climb_category"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	Country       string    `json:"country"`
	Private       bool      `json:"private"`
	Hazardous     bool      `json:"hazardous"`
	Starred       bool      `json:"starred"`
}

// DetailedActivity represents detailed information about an activity.
type DetailedActivity struct {
	ID                   uint                    `json:"id"`
	ExternalID           string                  `json:"external_id"`
	UploadID             uint                    `json:"upload_id"`
	Athlete              SummaryAthlete          `json:"athlete"`
	Name                 string                  `json:"name"`
	Distance             float64                 `json:"distance"`
	MovingTime           int                     `json:"moving_time"`
	ElapsedTime          int                     `json:"elapsed_time"`
	TotalElevationGain   float64                 `json:"total_elevation_gain"`
	Type                 string                  `json:"type"`
	SportType            string                  `json:"sport_type"`
	StartDate            time.Time               `json:"start_date"`
	StartDateLocal       time.Time               `json:"start_date_local"`
	Timezone             string                  `json:"timezone"`
	StartLatlng          []float64               `json:"start_latlng"`
	EndLatlng            []float64               `json:"end_latlng"`
	AchievementCount     int                     `json:"achievement_count"`
	KudosCount           int                     `json:"kudos_count"`
	CommentCount         int                     `json:"comment_count"`
	AthleteCount         int                     `json:"athlete_count"`
	PhotoCount           int                     `json:"photo_count"`
	Map                  PolylineMap             `json:"map"`
	Trainer              bool                    `json:"trainer"`
	Commute              bool                    `json:"commute"`
	Manual               bool                    `json:"manual"`
	Private              bool                    `json:"private"`
	Flagged              bool                    `json:"flagged"`
	WorkoutType          int                     `json:"workout_type"`
	UploadIDStr          string                  `json:"upload_id_str"`
	AverageSpeed         float64                 `json:"average_speed"`
	MaxSpeed             float64                 `json:"max_speed"`
	HasKudoed            bool                    `json:"has_kudoed"`
	GearID               string                  `json:"gear_id"`
	Kilojoules           float64                 `json:"kilojoules"`
	AverageWatts         float64                 `json:"average_watts"`
	DeviceWatts          bool                    `json:"device_watts"`
	MaxWatts             int                     `json:"max_watts"`
	WeightedAverageWatts int                     `json:"weighted_average_watts"`
	Description          string                  `json:"description"`
	Photos               PhotosSummary           `json:"photos"`
	Gear                 SummaryGear             `json:"gear"`
	Calories             float64                 `json:"calories"`
	SegmentEfforts       []DetailedSegmentEffort `json:"segment_efforts"`
	DeviceName           string                  `json:"device_name"`
	EmbedToken           string                  `json:"embed_token"`
	SplitsMetric         []Split                 `json:"splits_metric"`
	SplitsStandard       []Split                 `json:"splits_standard"`
	Laps                 []Lap                   `json:"laps"`
	BestEfforts          []DetailedSegmentEffort `json:"best_efforts"`
	AverageCadence       float64                 `json:"average_cadence"`
	AverageHeartrate     float64                 `json:"average_heartrate"`
}

// PhotosSummary represents a summary of photos for an activity.
type PhotosSummary struct {
	Count           int   `json:"count"`
	Primary         Photo `json:"primary"`
	UsePrimaryPhoto bool  `json:"use_primary_photo"`
}

// Photo represents a photo associated with an activity.
type Photo struct {
	ID       uint              `json:"id"`
	UniqueID string            `json:"unique_id"`
	URLs     map[string]string `json:"urls"`
	Source   int               `json:"source"`
}

// Split represents split information for an activity.
type Split struct {
	Distance            float64 `json:"distance"`
	ElapsedTime         int     `json:"elapsed_time"`
	ElevationDifference float64 `json:"elevation_difference"`
	MovingTime          int     `json:"moving_time"`
	Split               int     `json:"split"`
	AverageSpeed        float64 `json:"average_speed"`
	PaceZone            int     `json:"pace_zone"`
}

// Lap represents information about a lap in an activity.
type Lap struct {
	ID                 uint            `json:"id"`
	Activity           SummaryActivity `json:"activity"`
	Athlete            SummaryAthlete  `json:"athlete"`
	AverageCadence     float64         `json:"average_cadence"`
	AverageSpeed       float64         `json:"average_speed"`
	AverageHeartrate   float64         `json:"average_heartrate"`
	Distance           float64         `json:"distance"`
	ElapsedTime        int             `json:"elapsed_time"`
	StartIndex         int             `json:"start_index"`
	EndIndex           int             `json:"end_index"`
	LapIndex           int             `json:"lap_index"`
	MaxSpeed           float64         `json:"max_speed"`
	MaxHeartrate       float64         `json:"max_heartrate"`
	MovingTime         int             `json:"moving_time"`
	Name               string          `json:"name"`
	PaceZone           int             `json:"pace_zone"`
	Split              int             `json:"split"`
	StartDate          time.Time       `json:"start_date"`
	StartDateLocal     time.Time       `json:"start_date_local"`
	TotalElevationGain float64         `json:"total_elevation_gain"`
}

// SportType represents the type of sport/activity
type SportType string

const (
	SportTypeAlpineSki                     SportType = "AlpineSki"
	SportTypeBackcountrySki                SportType = "BackcountrySki"
	SportTypeBadminton                     SportType = "Badminton"
	SportTypeCanoeing                      SportType = "Canoeing"
	SportTypeCrossfit                      SportType = "Crossfit"
	SportTypeEBikeRide                     SportType = "EBikeRide"
	SportTypeElliptical                    SportType = "Elliptical"
	SportTypeEMountainBikeRide             SportType = "EMountainBikeRide"
	SportTypeGolf                          SportType = "Golf"
	SportTypeGravelRide                    SportType = "GravelRide"
	SportTypeHandcycle                     SportType = "Handcycle"
	SportTypeHighIntensityIntervalTraining SportType = "HighIntensityIntervalTraining"
	SportTypeHike                          SportType = "Hike"
	SportTypeIceSkate                      SportType = "IceSkate"
	SportTypeInlineSkate                   SportType = "InlineSkate"
	SportTypeKayaking                      SportType = "Kayaking"
	SportTypeKitesurf                      SportType = "Kitesurf"
	SportTypeMountainBikeRide              SportType = "MountainBikeRide"
	SportTypeNordicSki                     SportType = "NordicSki"
	SportTypePickleball                    SportType = "Pickleball"
	SportTypePilates                       SportType = "Pilates"
	SportTypeRacquetball                   SportType = "Racquetball"
	SportTypeRide                          SportType = "Ride"
	SportTypeRockClimbing                  SportType = "RockClimbing"
	SportTypeRollerSki                     SportType = "RollerSki"
	SportTypeRowing                        SportType = "Rowing"
	SportTypeRun                           SportType = "Run"
	SportTypeSail                          SportType = "Sail"
	SportTypeSkateboard                    SportType = "Skateboard"
	SportTypeSnowboard                     SportType = "Snowboard"
	SportTypeSnowshoe                      SportType = "Snowshoe"
	SportTypeSoccer                        SportType = "Soccer"
	SportTypeSquash                        SportType = "Squash"
	SportTypeStairStepper                  SportType = "StairStepper"
	SportTypeStandUpPaddling               SportType = "StandUpPaddling"
	SportTypeSurfing                       SportType = "Surfing"
	SportTypeSwim                          SportType = "Swim"
	SportTypeTableTennis                   SportType = "TableTennis"
	SportTypeTennis                        SportType = "Tennis"
	SportTypeTrailRun                      SportType = "TrailRun"
	SportTypeVelomobile                    SportType = "Velomobile"
	SportTypeVirtualRide                   SportType = "VirtualRide"
	SportTypeVirtualRow                    SportType = "VirtualRow"
	SportTypeVirtualRun                    SportType = "VirtualRun"
	SportTypeWalk                          SportType = "Walk"
	SportTypeWeightTraining                SportType = "WeightTraining"
	SportTypeWheelchair                    SportType = "Wheelchair"
	SportTypeWindsurf                      SportType = "Windsurf"
	SportTypeWorkout                       SportType = "Workout"
	SportTypeYoga                          SportType = "Yoga"
)

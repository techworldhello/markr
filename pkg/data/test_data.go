package data

import (
	"time"
)

var RequestBody = `<?xml version="1.0" encoding="UTF-8" ?>
<mcq-test-results>
	<mcq-test-result scanned-on="2017-12-04T12:12:10+11:00">
		<first-name>KJ</first-name>
		<last-name>Alysander</last-name>
		<student-number>002299</student-number>
		<test-id>9863</test-id>
		<answer question="0" marks-available="1" marks-awarded="1">D</answer>
		<answer question="1" marks-available="1" marks-awarded="1">D</answer>
		<answer question="2" marks-available="1" marks-awarded="1">D</answer>
		<answer question="3" marks-available="1" marks-awarded="0">C</answer>
		<answer question="4" marks-available="1" marks-awarded="1">B</answer>
		<answer question="5" marks-available="1" marks-awarded="0">D</answer>
		<answer question="6" marks-available="1" marks-awarded="0">A</answer>
		<answer question="7" marks-available="1" marks-awarded="1">A</answer>
		<answer question="8" marks-available="1" marks-awarded="1">B</answer>
		<answer question="9" marks-available="1" marks-awarded="1">D</answer>
		<answer question="10" marks-available="1" marks-awarded="1">A</answer>
		<answer question="11" marks-available="1" marks-awarded="1">B</answer>
		<answer question="12" marks-available="1" marks-awarded="0">A</answer>
		<answer question="13" marks-available="1" marks-awarded="0">B</answer>
		<answer question="14" marks-available="1" marks-awarded="1">B</answer>
		<answer question="15" marks-available="1" marks-awarded="1">A</answer>
		<answer question="16" marks-available="1" marks-awarded="1">C</answer>
		<answer question="17" marks-available="1" marks-awarded="0">B</answer>
		<answer question="18" marks-available="1" marks-awarded="1">A</answer>
		<answer question="19" marks-available="1" marks-awarded="0">B</answer>
		<summary-marks available="20" obtained="13" />
	</mcq-test-result>
	<mcq-test-result scanned-on="2017-12-04T12:13:10+11:00">
		<first-name>KJ</first-name>
		<last-name>Jim</last-name>
		<student-number>2300</student-number>
		<test-id>9863</test-id>
		<answer question="0" marks-available="1" marks-awarded="0">C</answer>
		<answer question="1" marks-available="1" marks-awarded="0">B</answer>
		<answer question="2" marks-available="1" marks-awarded="0">D</answer>
		<answer question="3" marks-available="1" marks-awarded="1">A</answer>
		<answer question="4" marks-available="1" marks-awarded="1">C</answer>
		<answer question="5" marks-available="1" marks-awarded="0">C</answer>
		<answer question="6" marks-available="1" marks-awarded="0">C</answer>
		<answer question="7" marks-available="1" marks-awarded="0">B</answer>
		<answer question="8" marks-available="1" marks-awarded="0">C</answer>
		<answer question="9" marks-available="1" marks-awarded="1">C</answer>
		<answer question="10" marks-available="1" marks-awarded="0">C</answer>
		<answer question="11" marks-available="1" marks-awarded="0">B</answer>
		<answer question="12" marks-available="1" marks-awarded="0">B</answer>
		<answer question="13" marks-available="1" marks-awarded="0">B</answer>
		<answer question="14" marks-available="1" marks-awarded="1">B</answer>
		<answer question="15" marks-available="1" marks-awarded="0">C</answer>
		<answer question="16" marks-available="1" marks-awarded="1">C</answer>
		<answer question="17" marks-available="1" marks-awarded="1">A</answer>
		<answer question="18" marks-available="1" marks-awarded="1">C</answer>
		<answer question="19" marks-available="1" marks-awarded="1">D</answer>
		<summary-marks available="20" obtained="8" />
	</mcq-test-result>
</mcq-test-results>`

var IncompleteRequestBody = `<?xml version="1.0" encoding="UTF-8" ?>
<mcq-test-results>
	<mcq-test-result scanned-on="2017-12-04T12:12:10+11:00">
		<first-name>KJ</first-name>
		<last-name>Alysander</last-name>
		<student-number>002299</student-number>
	</mcq-test-result>
</mcq-test-results>`

var ScannedTime = time.Now()

func GetTestResults() McqTestResults{
	return McqTestResults{
		Results: []*TestResult{
			{
				ScannedOn: ScannedTime,
				FirstName: "Daniel",
				LastName: "Craig",
				StudentNumber: 1234,
				TestID: 1,
				SummaryMarks: &SummaryMarks{
					Available: 20,
					Obtained: 18,
				},
			},
			{
				ScannedOn: ScannedTime,
				FirstName: "Jane",
				LastName: "Doe",
				StudentNumber: 4321,
				TestID: 2,
				SummaryMarks: &SummaryMarks{
					Available: 20,
					Obtained: 10,
				},
			},
			{
				ScannedOn: ScannedTime,
				FirstName: "Spongebob",
				LastName: "Squarepants",
				StudentNumber: 1212,
				TestID: 3,
				SummaryMarks: &SummaryMarks{
					Available: 20,
					Obtained: 14,
				},
			},
		},
	}
}

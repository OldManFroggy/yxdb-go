package spatial_test

import (
	"fmt"
	"github.com/tlarsendataguy-yxdb/yxdb-go"
	"github.com/tlarsendataguy-yxdb/yxdb-go/spatial"
	"strings"
	"testing"
)

func TestPoint(t *testing.T) {
	expected := `{"type": "Point", "coordinates": [-96.679688, 37.230328]}`
	path := `../test_files/point.yxdb`
	err := testSpatial(path, expected)
	if err != nil {
		t.Fatalf(`got error: %v`, err.Error())
	}
}

func TestPoints(t *testing.T) {
	expected := `{"type": "MultiPoint", "coordinates": [[-113.730469, 7.885147], [-113.378906, 46.679594], [-100.019531, 40.178873], [-88.769531, 49.61071], [-85.957031, 12.039321]]}`
	path := `../test_files/multi-point.yxdb`
	err := testSpatial(path, expected)
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}
}

func TestLine(t *testing.T) {
	expected := `{"type": "LineString", "coordinates": [[-106.875, 42.293564], [-84.375, 41.244772], [-106.347656, 36.738884], [-85.253906, 35.173808], [-110.390625, 32.546813], [-89.472656, 29.22889]]}`
	path := `../test_files/line.yxdb`
	err := testSpatial(path, expected)
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}
}

func TestLines(t *testing.T) {
	expected := `{"type": "MultiLineString", "coordinates": [[[-92.285156, 55.875311], [-74.355469, 53.225768], [-76.992188, 41.902277], [-76.992188, 29.382175], [-66.269531, 43.068888]], [[-108.984375, 43.197167], [-70.664063, 49.037868], [-97.558594, 25.799891], [-73.125, 21.616579], [-97.910156, 4.565474], [-81.386719, -3.513421]], [[-121.464844, 45.213004], [-109.6875, -0.175781]], [[-114.082031, 57.231503], [-107.753906, 55.677584], [-111.972656, 51.399206], [-120.9375, 54.470038], [-122.34375, 58.995311], [-115.136719, 62.103883], [-104.0625, 59.085739], [-101.777344, 51.944265], [-108.28125, 47.517201], [-123.222656, 50.176898]]]}`
	path := `../test_files/multi-line.yxdb`
	err := testSpatial(path, expected)
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}
}

func TestPoly(t *testing.T) {
	expected := `{"type": "Polygon", "coordinates": [[[-84.550781, 42.811522], [-101.25, 34.452218], [-96.855469, 43.068888], [-114.082031, 32.10119], [-119.355469, 42.55308], [-104.589844, 44.087585], [-107.402344, 47.279229], [-91.230469, 52.908902], [-84.550781, 42.811522]]]}`
	path := `../test_files/poly.yxdb`
	err := testSpatial(path, expected)
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}
}

func TestPolys(t *testing.T) {
	expected := `{"type": "MultiPolygon", "coordinates": [[[[-107.226562, 55.973798], [-109.335938, 52.696361], [-114.257813, 55.578345], [-107.226562, 55.973798]], [[-89.824219, 42.811522], [-97.382813, 36.597889], [-106.347656, 40.84706], [-105.46875, 47.040182], [-96.328125, 46.437857], [-89.824219, 42.811522]], [[-71.542969, 36.879621], [-74.53125, 31.802893], [-88.769531, 26.902477], [-91.933594, 33.870416], [-86.484375, 38.822591], [-76.289063, 40.313043], [-71.542969, 36.879621]], [[-68.027344, 52.802761], [-70.3125, 49.267805], [-80.332031, 48.224673], [-86.835938, 49.382373], [-91.054688, 53.120405], [-88.417969, 56.752723], [-84.550781, 54.775346], [-73.125, 55.37911], [-68.027344, 52.802761]]]]}`
	path := `../test_files/multi-poly.yxdb`
	err := testSpatial(path, expected)
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}
}

func TestPolyWithHole(t *testing.T) {
	expected := `{"type": "MultiPolygon", "coordinates": [[[[-88.417969, 41.508577], [-89.121094, 16.299051], [-106.875, 15.961329], [-106.171875, 42.811522], [-88.417969, 41.508577]], [[-78.75, 47.872144], [-114.257813, 47.279229], [-115.3125, 8.581021], [-80.15625, 9.102097], [-78.75, 47.872144]], [[-68.203125, 54.673831], [-70.664063, 1.406109], [-124.628906, 0.35156], [-123.574219, 53.014783], [-68.203125, 54.673831]]]]}`
	path := `../test_files/multi-poly-holes.yxdb`
	err := testSpatial(path, expected)
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}
}

func TestNullSpatial(t *testing.T) {
	expected := ``
	path := `../test_files/null-spatial.yxdb`
	err := testSpatial(path, expected)
	if err != nil {
		t.Fatalf(`got error %v`, err.Error())
	}
}

func TestInvalidObjectType(t *testing.T) {
	data := []byte{1, 0, 0, 0, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34, 34}
	_, err := spatial.ToGeoJSON(data)
	if err == nil {
		t.Fatalf(`expected an error but got none`)
	}
	t.Logf(`error: %v`, err.Error())
}

func TestFileTooShort(t *testing.T) {
	data := []byte{1, 0, 0, 0}
	_, err := spatial.ToGeoJSON(data)
	if err == nil {
		t.Fatalf(`expected an error but got none`)
	}
	t.Logf(`error: %v`, err.Error())
}

func testSpatial(path string, expected string) error {
	expected = strings.Replace(expected, ` `, ``, -1)
	reader, err := yxdb.ReadFile(path)
	if err != nil {
		return err
	}
	for reader.Next() {
		blob := reader.ReadBlobWithIndex(1)
		geo, err := spatial.ToGeoJSON(blob)
		if err != nil {
			return err
		}
		if expected != geo {
			return fmt.Errorf("expected\n%v\nbut got\n%v", expected, geo)
		}
	}
	return nil
}

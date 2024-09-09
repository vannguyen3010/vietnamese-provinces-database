package dumper

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestGetAdministrativeUnit_ProvinceLevel(t *testing.T) {
	assert.Equal(t, 1, getAdministrativeUnit_ProvinceLevel("Thành phố Hồ Chí Minh"))
	assert.Equal(t, 2, getAdministrativeUnit_ProvinceLevel("Tỉnh Khánh Hoà"))
	assert.Equal(t, 2, getAdministrativeUnit_ProvinceLevel("Tỉnh Thành phố"))
}

func TestGetAdministrativeUnit_ProvinceLevel_ShouldThrowException(t *testing.T) {
	assert.PanicsWithValue(
		t, "Unable to determine administrative unit name from province: Quận Ba Đình",
		func() {
			getAdministrativeUnit_ProvinceLevel("Quận Ba Đình")
	})
}

func TestGetAdministrativeUnit_DistrictLevel(t *testing.T) {
	assert.Equal(t, 4, getAdministrativeUnit_DistrictLevel("Thành phố Cao Bằng"))

	assert.Equal(t, 5, getAdministrativeUnit_DistrictLevel("Quận 2"))
	assert.Equal(t, 6, getAdministrativeUnit_DistrictLevel("Thị xã Đông Hòa"))
	assert.Equal(t, 6, getAdministrativeUnit_DistrictLevel("Thị xã Quận Thành"))

	assert.Equal(t, 7, getAdministrativeUnit_DistrictLevel("Huyện Chợ Lách"))
	assert.Equal(t, 7, getAdministrativeUnit_DistrictLevel("Huyện Thành phố Quận"))

}

func TestGetAdministrativeUnit_DistrictLevel_SpecialCase(t *testing.T) {
	// special case
	assert.Equal(t, 3, getAdministrativeUnit_DistrictLevel("Thành phố Thủ Đức"))
}

func TestGetAdministrativeUnit_DistrictLevel_Exception(t *testing.T) {
	assert.PanicsWithValue(
		t, "Unable to determine administrative unit name from district: Tỉnh Hà Giang",
		func() {
			getAdministrativeUnit_DistrictLevel("Tỉnh Hà Giang")
	})
}

func TestGetAdministrativeUnit_WardLevel(t *testing.T) {
	assert.Equal(t, 8, getAdministrativeUnit_WardLevel("Phường Phường Đúc"))
	assert.Equal(t, 9, getAdministrativeUnit_WardLevel("Thị trấn Châu Thành"))
	assert.Equal(t, 10, getAdministrativeUnit_WardLevel("Xã Tân Xã"))
}

func TestGetAdministrativeUnit_WardLevel_Exception(t *testing.T) {
	assert.PanicsWithValue(
		t, "Unable to determine administrative unit name from ward: Quận 9",
		func() {
			getAdministrativeUnit_WardLevel("Quận 9")
	})
}

// Test string normalization

func TestNormalizeString(t *testing.T) {
	assert.Equal(t, "Da Lat", normalizeString("Đà Lạt"))
	assert.Equal(t, "Hoi An", normalizeString("Hội An"))
	assert.Equal(t, "Bai bien Cua Lo", normalizeString("Bãi biển Cửa Lò"))
	assert.Equal(t, "Ghenh da Sen Thang Bang", normalizeString("Ghềnh đá Sên Thắng Bâng"))
	assert.Equal(t, "Ong Ot Enh Uong", normalizeString("Ông Ớt Ễnh Ương"))
}

func TestToCodeName(t *testing.T) {
	assert.Equal(t, "phong_nha_ke_bang", toCodeName("Phong Nha - Kẻ Bàng"))
	assert.Equal(t, "nieng", toCodeName("N'Iêng"))
	assert.Equal(t, "thanh_hoa", toCodeName("Thanh Hoá"))
}

func TestRemoveWhiteSpaces(t *testing.T) {
	assert.Equal(t, "Ninh Binh", removeWhiteSpaces(" Ninh  Binh "))
	assert.Equal(t, "Quang Ninh", removeWhiteSpaces(" Quang Ninh"))
	assert.Equal(t, "Vinh Phuc", removeWhiteSpaces("Vinh Phuc "))
}



func TestSanitizeString(t *testing.T) {
	
}

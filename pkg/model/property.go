package model

import (
	"next-terminal/pkg/global"
	"next-terminal/pkg/guacd"
)

const (
	SshMode = "ssh-mode"
)

type Property struct {
	Name  string `gorm:"primary_key" json:"name"`
	Value string `json:"value"`
}

func (r *Property) TableName() string {
	return "properties"
}

func FindAllProperties() (o []Property) {
	if global.DB.Find(&o).Error != nil {
		return nil
	}
	return
}

func CreateNewProperty(o *Property) (err error) {
	err = global.DB.Create(o).Error
	return
}

func UpdatePropertyByName(o *Property, name string) {
	o.Name = name
	global.DB.Updates(o)
}

func FindPropertyByName(name string) (o Property, err error) {
	err = global.DB.Where("name = ?", name).First(&o).Error
	return
}

func FindAllPropertiesMap() map[string]string {
	properties := FindAllProperties()
	propertyMap := make(map[string]string)
	for i := range properties {
		propertyMap[properties[i].Name] = properties[i].Value
	}
	return propertyMap
}

func GetDrivePath() (string, error) {
	property, err := FindPropertyByName(guacd.DrivePath)
	if err != nil {
		return "", err
	}
	return property.Value, nil
}

func GetRecordingPath() (string, error) {
	property, err := FindPropertyByName(guacd.RecordingPath)
	if err != nil {
		return "", err
	}
	return property.Value, nil
}

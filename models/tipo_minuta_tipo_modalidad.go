package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type TipoMinutaTipoModalidad struct {
<<<<<<< HEAD
	Id              int            `orm:"column(id);auto"`
	TipoMinutaId    *TipoMinuta    `orm:"column(tipo_minuta_id);rel(fk)"`
	TipoModalidadId *TipoModalidad `orm:"column(tipo_modalidad_id);rel(fk)"`
=======
	Id              int `orm:"column(id);auto"`
	TipoMinutaId    int `orm:"column(tipo_minuta_id);null"`
	TipoModalidadId int `orm:"column(tipo_modalidad_id);null"`
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
}

func (t *TipoMinutaTipoModalidad) TableName() string {
	return "tipo_minuta_tipo_modalidad"
}

func init() {
	orm.RegisterModel(new(TipoMinutaTipoModalidad))
}

// AddTipoMinutaTipoModalidad insert a new TipoMinutaTipoModalidad into database and returns
// last inserted Id on success.
func AddTipoMinutaTipoModalidad(m *TipoMinutaTipoModalidad) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTipoMinutaTipoModalidadById retrieves TipoMinutaTipoModalidad by Id. Returns error if
// Id doesn't exist
func GetTipoMinutaTipoModalidadById(id int) (v *TipoMinutaTipoModalidad, err error) {
	o := orm.NewOrm()
	v = &TipoMinutaTipoModalidad{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTipoMinutaTipoModalidad retrieves all TipoMinutaTipoModalidad matches certain condition. Returns empty list if
// no records exist
func GetAllTipoMinutaTipoModalidad(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(TipoMinutaTipoModalidad))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
<<<<<<< HEAD
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
=======
		qs = qs.Filter(k, v)
>>>>>>> c1d187705a12d73b3be4fd94851e77702888b270
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []TipoMinutaTipoModalidad
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTipoMinutaTipoModalidad updates TipoMinutaTipoModalidad by Id and returns error if
// the record to be updated doesn't exist
func UpdateTipoMinutaTipoModalidadById(m *TipoMinutaTipoModalidad) (err error) {
	o := orm.NewOrm()
	v := TipoMinutaTipoModalidad{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTipoMinutaTipoModalidad deletes TipoMinutaTipoModalidad by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTipoMinutaTipoModalidad(id int) (err error) {
	o := orm.NewOrm()
	v := TipoMinutaTipoModalidad{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&TipoMinutaTipoModalidad{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

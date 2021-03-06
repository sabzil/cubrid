package gci

/*
#include <stdio.h>
#include <stdlib.h>
#include "cas_cci.h"
#include "cas_error.h"
int ex_cci_connect(char *ip, int port, char *db_name, char *db_user, char *db_password) {
	int con = cci_connect(ip, port, db_name, db_user, db_password);
	return con;
}

char* ex_cci_get_result_info_name(T_CCI_COL_INFO* res_info, int index) {
	return CCI_GET_RESULT_INFO_NAME(res_info, index);
}

T_CCI_U_TYPE ex_cci_get_result_info_type(T_CCI_COL_INFO* res_info, int index) {
	return CCI_GET_RESULT_INFO_TYPE(res_info, index);
}

int ex_cci_is_set_type(T_CCI_U_TYPE type) {
	return CCI_IS_SET_TYPE(type);
}

int ex_cci_is_collection_type(T_CCI_U_TYPE type) {
	return CCI_IS_COLLECTION_TYPE(type);
}
*/
import "C"
import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"unsafe"
)

func Bind_param() int {
	//cci_bind_param
	return 0
}

func Bind_param_int(req_handle int, index int, value interface{}, flag int) int {
	var handle C.int = C.int(req_handle)
	var res C.int

	c_param := C.int(value.(int64))
	res = C.cci_bind_param(handle, C.int(index), C.CCI_A_TYPE_INT,
		unsafe.Pointer(&c_param), C.CCI_U_TYPE_INT, C.char(flag))

	return int(res)
}

func Bind_param_string(req_handle int, index int, value interface{}, flag int) int {
	var handle C.int = C.int(req_handle)
	var res C.int

	ss := fmt.Sprint(value)
	res = C.cci_bind_param(handle, C.int(index), C.CCI_A_TYPE_STR,
		unsafe.Pointer(C.CString(ss)), C.CCI_U_TYPE_STRING, C.char(flag))

	return int(res)
}

func Bind_param_float(req_handle int, index int, value interface{}, flag int) int {
	var handle C.int = C.int(req_handle)
	var res C.int

	c_param := C.float(value.(float64))
	res = C.cci_bind_param(handle, C.int(index), C.CCI_A_TYPE_FLOAT,
		unsafe.Pointer(&c_param), C.CCI_U_TYPE_FLOAT, C.char(flag))

	return int(res)
}

func Bind_param_array() int {
	//cci_bind_param_array
	return 0
}

func Bind_param_array_size(req_handle int, array_size int) int {
	var cHandle C.int = C.int(req_handle)
	var cSize C.int = C.int(array_size)

	res := C.cci_bind_param_array_size(cHandle, cSize)

	return int(res)
}

func Blob_free(blob GCI_BLOB) {
	var data C.T_CCI_BLOB = C.T_CCI_BLOB(blob)
	C.cci_blob_free(data)
}

func Blob_new(conn_handle int) (GCI_BLOB, GCI_ERROR) {
	var err GCI_ERROR

	return nil, err
}

func Blob_read(con_handle int, blob GCI_BLOB, start_pos int64, length int64) (GCI_BLOB, GCI_ERROR) {
	var handle C.int = C.int(con_handle)
	var res C.int
	var c_start_pos C.longlong = C.longlong(start_pos)
	var c_length C.int = C.int(length)
	var c_blob string
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR
	var data C.T_CCI_BLOB = C.T_CCI_BLOB(blob)
	var res_blob GCI_BLOB

	c_buf := C.CString(c_blob)
	defer C.free(unsafe.Pointer(c_buf))
	res = C.cci_blob_read(handle, data, c_start_pos, c_length, c_buf, &cci_error)
	if res < C.int(0) {
		err.Code = int(cci_error.err_code)
		err.Msg = C.GoString(&cci_error.err_msg[0])
	}

	res_blob = GCI_BLOB(c_buf)

	return res_blob, err
}

func Blob_size(blob GCI_BLOB) int64 {
	var size C.longlong
	var data C.T_CCI_BLOB = C.T_CCI_BLOB(blob)

	size = C.cci_blob_size(data)

	return int64(size)
}

func Blob_write() {

}

func Cancel(conn_handle int) int {
	var cHandle C.int = C.int(conn_handle)

	res := C.cci_cancel(cHandle)

	return int(res)
}

func Clob_free() {

}

func Clob_new() {

}

func Clob_read() {

}

func Clob_size() {

}

func Clob_write() {

}

func Close_query_result() {

}

func Close_req_handle(req_handle int) int {
	var err C.int
	var handle C.int = C.int(req_handle)

	err = C.cci_close_req_handle(handle)

	return int(err)
}

func Col_get() {

}

func Col_seq_drop() {

}

func Col_seq_insert() {

}

func Col_seq_put() {

}

func Col_set_add() {

}

func Col_set_drop() {

}

func Col_size() {

}

func Connect(ip string, port int, db_name string, db_user string, db_password string) int {
	serverAddress := C.CString(ip)
	serverPort := C.int(port)
	dbName := C.CString(db_name)
	dbUser := C.CString(db_user)
	dbPassword := C.CString(db_password)

	defer C.free(unsafe.Pointer(serverAddress))
	defer C.free(unsafe.Pointer(dbName))
	defer C.free(unsafe.Pointer(dbUser))
	defer C.free(unsafe.Pointer(dbPassword))

	con := C.ex_cci_connect(serverAddress, serverPort, dbName, dbUser, dbPassword)
	return int(con)
}

func Connect_with_url(url string, user string, password string) int {
	serverUrl := C.CString(url)
	serverUser := C.CString(user)
	serverPassword := C.CString(password)

	defer C.free(unsafe.Pointer(serverUrl))
	defer C.free(unsafe.Pointer(serverUser))
	defer C.free(unsafe.Pointer(serverPassword))

	con := C.cci_connect_with_url(serverUrl, serverUser, serverPassword)

	return int(con)
}

func Cursor(req_handle int, offset int, origin GCI_CURSOR_POS) (int, GCI_ERROR) {
	var handle C.int
	var c_offset C.int
	var c_origin C.T_CCI_CURSOR_POS
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR
	var res C.int

	handle = C.int(req_handle)
	c_offset = C.int(offset)
	c_origin = C.T_CCI_CURSOR_POS(origin)

	res = C.cci_cursor(handle, c_offset, c_origin, &cci_error)
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])

	return int(res), err
}

func Cursor_update() {

}

func Datasource_borrow() {

}

func Datasource_change_property() {

}

func Datasource_create() {

}

func Datasource_destroy() {

}

func Datasource_release() {

}

func Disconnect(conn_handle int) (int, GCI_ERROR) {
	var cHandle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var res C.int
	var err GCI_ERROR

	res = C.cci_disconnect(cHandle, &cci_error)
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])

	return int(res), err
}

func End() {
	C.cci_end()
}

func End_tran(conn_handle int, tran_type int) (int, GCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR

	res = C.cci_end_tran(handle, C.char(tran_type), &cci_error)
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])

	return int(res), err
}

func Escape_string() {

}

func Execute(req_handle int, flag int, max_col_size int) (int, GCI_ERROR) {
	var res C.int
	var cci_error C.T_CCI_ERROR
	var handle C.int = C.int(req_handle)
	var err GCI_ERROR

	res = C.cci_execute(handle, C.char(flag), C.int(max_col_size), &cci_error)
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])

	return int(res), err
}

func Execute_array() {

}

func Execute_batch() {

}

func Execute_result() {

}

func Fetch(req_handle int) (int, GCI_ERROR) {
	var handle C.int = C.int(req_handle)
	var cci_error C.T_CCI_ERROR
	var res C.int
	var err GCI_ERROR

	res = C.cci_fetch(handle, &cci_error)
	if res < C.int(0) {
		err.Code = int(cci_error.err_code)
		err.Msg = C.GoString(&cci_error.err_msg[0])
	}

	return int(res), err
}

func Fetch_buffer_clear() {

}

func Fetch_sensitive() {

}

func Fetch_size() {

}

func Get_autocommit() {

}

func Get_bind_num(req_handle int) int {
	var param_cnt C.int
	var handle C.int = C.int(req_handle)

	param_cnt = C.cci_get_bind_num(handle)

	return int(param_cnt)
}

func Get_cas_info() {

}

func Get_class_num_objs() {

}

func Get_cur_oid() {

}

func Get_data() {

}

func Get_data_string(req_handle int, idx int) (int, string, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf *C.char
	var res C.int
	var indicator C.int
	var data string

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_STR, unsafe.Pointer(&buf), &indicator)
	data = C.GoString(buf)

	return int(res), data, int(indicator)
}

func Get_data_int(req_handle int, idx int) (int, int, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf C.int
	var res C.int
	var indicator C.int
	var data int

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_INT, unsafe.Pointer(&buf), &indicator)
	data = int(buf)

	return int(res), data, int(indicator)
}

func Get_data_float(req_handle int, idx int) (int, float64, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf C.float
	var res C.int
	var indicator C.int
	var data float64

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_FLOAT, unsafe.Pointer(&buf), &indicator)
	data = float64(buf)

	return int(res), data, int(indicator)
}

func Get_data_double(req_handle int, idx int) (int, float64, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf C.double
	var res C.int
	var indicator C.int
	var data float64

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_DOUBLE, unsafe.Pointer(&buf), &indicator)
	data = float64(buf)

	return int(res), data, int(indicator)
}

func Get_data_bit(req_handle int, idx int) (int, GCI_BIT, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf C.T_CCI_BIT
	var res C.int
	var indicator C.int
	var data GCI_BIT

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_BIT, unsafe.Pointer(&buf), &indicator)
	data.size = int(buf.size)
	data.buf = C.GoBytes(unsafe.Pointer(buf.buf), buf.size)

	return int(res), data, int(indicator)
}

func Get_data_set(req_handle int, idx int) (int, GCI_SET, int) {
	var handle C.int = C.int(req_handle)
	var c_idx C.int = C.int(idx)
	var buf C.T_CCI_SET
	var res C.int
	var indicator C.int
	var data GCI_SET

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_SET, unsafe.Pointer(&buf), &indicator)
	data = GCI_SET(buf)

	return int(res), data, int(indicator)
}

func Get_data_date(req_handle int, idx int) (int, GCI_DATE, int) {
	log.Println("gci_get_data_date_start")
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.T_CCI_DATE
	var res C.int
	var indicator C.int
	var data GCI_DATE

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_DATE,
		unsafe.Pointer(&buf), &indicator)
	data.yr = int(buf.yr)
	data.mon = int(buf.mon)
	data.day = int(buf.day)
	data.hh = int(buf.hh)
	data.mm = int(buf.mm)
	data.ss = int(buf.ss)
	data.ms = int(buf.ms)

	log.Println("gci_get_data_date_end")
	return int(res), data, int(indicator)
}

func Get_data_bigint(req_handle int, idx int) (int, int64, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.int64_t
	var res C.int
	var indicator C.int
	var data int64

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_BIGINT,
		unsafe.Pointer(&buf), &indicator)
	data = int64(buf)

	return int(res), data, int(indicator)
}

func Get_data_blob(req_handle int, idx int) (int, GCI_BLOB, int) {
	var handle C.int = C.int(req_handle)
	var c_idx = C.int(idx)
	var buf C.T_CCI_BLOB
	var res C.int
	var indicator C.int
	var data GCI_BLOB

	res = C.cci_get_data(handle, c_idx, C.CCI_A_TYPE_BLOB,
		unsafe.Pointer(&buf), &indicator)
	data = GCI_BLOB(buf)

	return int(res), data, int(indicator)
}

func Get_db_parameter() {

}

func Get_db_version() {

}

func Get_err_msg() {

}

func Get_error_msg() {

}

func Get_holdability() {

}

func Get_last_insert_id(conn_handle int) (int64, GCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR
	var value *C.char
	var nid int64

	res = C.cci_get_last_insert_id(handle, unsafe.Pointer(value), &cci_error)
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])
	if res < 0 {
		return int64(res), err
	}

	id := C.GoString(value)
	nid, _ = strconv.ParseInt(id, 0, 64)

	return nid, err
}

func Get_login_timeout() {

}

func Get_query_plan() {

}

func Query_info_free() {

}

func Get_query_timeout() {

}

func Get_result_info(req_handle int) ([]GCI_COL_INFO, GCI_CUBRID_STMT, int) {
	var handle C.int = C.int(req_handle)
	var c_col_info *C.T_CCI_COL_INFO
	var go_col_info []C.T_CCI_COL_INFO
	var cubrid_stmt C.T_CCI_CUBRID_STMT
	var col_count C.int
	var gci_col_info []GCI_COL_INFO
	var gci_cubrid_stmt GCI_CUBRID_STMT

	c_col_info = C.cci_get_result_info(handle, &cubrid_stmt, &col_count)
	gci_cubrid_stmt = GCI_CUBRID_STMT(cubrid_stmt)

	sliceHeader := (*reflect.SliceHeader)((unsafe.Pointer(&go_col_info)))
	sliceHeader.Cap = int(col_count)
	sliceHeader.Len = int(col_count)
	sliceHeader.Data = uintptr(unsafe.Pointer(c_col_info))
	gci_col_info = make([]GCI_COL_INFO, int(col_count))
	for i := 0; i < int(col_count); i++ {
		gci_col_info[i].u_type = GCI_U_TYPE(go_col_info[C.int(i)]._type)
		gci_col_info[i].is_non_null = C.GoString(&go_col_info[C.int(i)].is_non_null)
		gci_col_info[i].scale = int16(go_col_info[C.int(i)].scale)
		gci_col_info[i].precision = int(go_col_info[C.int(i)].precision)
	}

	return gci_col_info, gci_cubrid_stmt, int(col_count)
}

func Get_result_info_name(col_info []GCI_COL_INFO, idx int) string {
	var result string
	result = col_info[idx-1].col_name
	return result
}

func Get_result_info_type(col_info []GCI_COL_INFO, idx int) GCI_U_TYPE {
	var result GCI_U_TYPE
	result = col_info[idx-1].u_type
	return result
}

func Get_version() {
}


func Init() {
	C.cci_init()
}

func Is_collection_type(u_type GCI_U_TYPE) int {
	var result int
	// 이게 맞는건가????
	res := (u_type) & GCI_CODE_COLLECTION
	if (res != 0) || ((u_type) == U_TYPE_SET) || ((u_type) == U_TYPE_MULTISET) || ((u_type) == U_TYPE_SEQUENCE) {
		result = 1
	} else {
		result = 0
	}
	return result
}


func Is_holdable() {

}

func Is_updatable() {

}

func Next_result() {

}

func Oid() {

}

func Oid_get() {

}

func Oid_get_class_name() {

}

func Oid_put() {

}

func Oid_put2() {

}

func Prepare(conn_handle int, sql_stmt string, flag byte) (int, GCI_ERROR) {
	var cHandle C.int = C.int(conn_handle)
	var cQuery *C.char = C.CString(sql_stmt)
	var cci_error C.T_CCI_ERROR
	var req C.int
	var err GCI_ERROR

	defer C.free(unsafe.Pointer(cQuery))

	req = C.cci_prepare(cHandle, cQuery, 0, &cci_error)
	err.Code = int(cci_error.err_code)
	err.Msg = C.GoString(&cci_error.err_msg[0])

	return int(req), err
}

func Prepare_and_execute() {

}

func Property_create() {

}

func Property_destroy() {

}

func Property_get() {

}

func Property_set() {

}

func Query_result_free() {

}


func Register_out_param() {

}

func Row_count(conn_handle int) (int64, GCI_ERROR) {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var row_count C.int
	var cci_error C.T_CCI_ERROR
	var err GCI_ERROR

	res = C.cci_row_count(handle, &row_count, &cci_error)
	if res < 0 {
		err.Code = int(cci_error.err_code)
		err.Msg = C.GoString(&cci_error.err_msg[0])
	}

	return int64(row_count), err
}

func Savepoint() {

}

func Schema_info() {

}

func Set_allocators() {

}

func Set_autocommit(conn_handle int, autocommit_mode AUTOCOMMIT_MODE) int {
	var res C.int
	var handle C.int = C.int(conn_handle)
	var mode C.int = C.int(autocommit_mode)

	res = C.cci_set_autocommit(handle, C.CCI_AUTOCOMMIT_MODE(mode))

	return int(res)
}

func Set_db_parameter() {

}

func Set_element_type() {

}

func Set_free(set GCI_SET) {
	var data C.T_CCI_SET = C.T_CCI_SET(set)

	C.cci_set_free(data)
}

/*
	set 버퍼 안에서의 index와 a_type
	ex) {'a', 'b', 'c'}
*/
func Set_get(set GCI_SET, index int, a_type GCI_A_TYPE) (int, interface{}, int) {
	var indicator int
	var res int
	var data interface{}

	switch a_type {
	case A_TYPE_STR:
		res, data, indicator = set_get_str(set, index)
	case A_TYPE_INT:
		res, data, indicator = set_get_int(set, index)
	case A_TYPE_FLOAT:
		res, data, indicator = set_get_float(set, index)
	case A_TYPE_DOUBLE:
		res, data, indicator = set_get_float(set, index)
	case A_TYPE_BIT:
		res, data, indicator = set_get_bit(set, index)
	case A_TYPE_DATE:
		res, data, indicator = set_get_date(set, index)
	case A_TYPE_BIGINT:
		res, data, indicator = set_get_bigint(set, index)
		// todo
		//case A_TYPE_BLOB:
		//	res, data, indicator = set_get_blob(set, index)
		//case A_TYPE_CLOB:
		//	res, data, indicator = set_get_clob(set, index)
	}

	return res, data, indicator
}

func Set_holdability() {

}

func Set_isolation_level() {

}

func Set_lock_timeout() {

}

func Set_login_timeout() {

}

func Set_make() {

}

func Set_max_row() {

}

func Set_query_timeout() {

}

func Set_size(set GCI_SET) int {
	var data C.T_CCI_SET = C.T_CCI_SET(set)
	var res C.int

	res = C.cci_set_size(data)

	return int(res)
}

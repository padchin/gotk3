// Same copyright and license as the rest of the files in this project
// This file contains accelerator related functions and structures

package gtk

// #include <gtk/gtk.h>
// #include "gtk.go.h"
import "C"
import (
	"unsafe"

	"github.com/d2r2/gotk3/gdk"
	"github.com/d2r2/gotk3/glib"
)

// AccelFlags is a representation of GTK's GtkAccelFlags
type AccelFlags int

const (
	ACCEL_VISIBLE AccelFlags = C.GTK_ACCEL_VISIBLE
	ACCEL_LOCKED  AccelFlags = C.GTK_ACCEL_LOCKED
	ACCEL_MASK    AccelFlags = C.GTK_ACCEL_MASK
)

func marshalAccelFlags(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum(C.toGValue(unsafe.Pointer(p)))
	return AccelFlags(c), nil
}

// AcceleratorName is a wrapper around gtk_accelerator_name().
func AcceleratorName(key uint, mods gdk.ModifierType) string {
	c := C.gtk_accelerator_name(C.guint(key), C.GdkModifierType(mods))
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// AcceleratorValid is a wrapper around gtk_accelerator_valid().
func AcceleratorValid(key uint, mods gdk.ModifierType) bool {
	return gobool(C.gtk_accelerator_valid(C.guint(key), C.GdkModifierType(mods)))
}

// AcceleratorGetDefaultModMask is a wrapper around gtk_accelerator_get_default_mod_mask().
func AcceleratorGetDefaultModMask() gdk.ModifierType {
	return gdk.ModifierType(C.gtk_accelerator_get_default_mod_mask())
}

// AcceleratorParse is a wrapper around gtk_accelerator_parse().
func AcceleratorParse(acc string) (key uint, mods gdk.ModifierType) {
	cstr := C.CString(acc)
	defer C.free(unsafe.Pointer(cstr))

	k := C.guint(0)
	m := C.GdkModifierType(0)

	C.gtk_accelerator_parse((*C.gchar)(cstr), &k, &m)
	return uint(k), gdk.ModifierType(m)
}

// AcceleratorGetLabel is a wrapper around gtk_accelerator_get_label().
func AcceleratorGetLabel(key uint, mods gdk.ModifierType) string {
	c := C.gtk_accelerator_get_label(C.guint(key), C.GdkModifierType(mods))
	defer C.g_free(C.gpointer(c))
	return goString(c)
}

// AcceleratorSetDefaultModMask is a wrapper around gtk_accelerator_set_default_mod_mask().
func AcceleratorSetDefaultModMask(mods gdk.ModifierType) {
	C.gtk_accelerator_set_default_mod_mask(C.GdkModifierType(mods))
}

/*
 * GtkAccelGroup
 */

// AccelGroup is a representation of GTK's GtkAccelGroup.
type AccelGroup struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkAccelGroup.
func (v *AccelGroup) native() *C.GtkAccelGroup {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkAccelGroup(ptr)
}

func marshalAccelGroup(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAccelGroup(obj), nil
}

func wrapAccelGroup(obj *glib.Object) *AccelGroup {
	return &AccelGroup{obj}
}

// AccelGroupNew is a wrapper around gtk_accel_group_new().
func AccelGroupNew() (*AccelGroup, error) {
	c := C.gtk_accel_group_new()
	if c == nil {
		return nil, nilPtrErr
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAccelGroup(obj), nil
}

// Connect is a wrapper around gtk_accel_group_connect().
func (v *AccelGroup) Connect(key uint, mods gdk.ModifierType, flags AccelFlags, f interface{}) {
	closure, _ := glib.ClosureNew(f)
	cl := (*C.GClosure)(unsafe.Pointer(closure))
	C.gtk_accel_group_connect(
		v.native(),
		C.guint(key),
		C.GdkModifierType(mods),
		C.GtkAccelFlags(flags),
		cl)
}

// ConnectByPath is a wrapper around gtk_accel_group_connect_by_path().
func (v *AccelGroup) ConnectByPath(path string, f interface{}) {
	closure, _ := glib.ClosureNew(f)
	cl := (*C.GClosure)(unsafe.Pointer(closure))

	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_group_connect_by_path(v.native(), (*C.gchar)(cstr), cl)
}

// Disconnect is a wrapper around gtk_accel_group_disconnect().
func (v *AccelGroup) Disconnect(f interface{}) {
	closure, _ := glib.ClosureNew(f)
	cl := (*C.GClosure)(unsafe.Pointer(closure))
	C.gtk_accel_group_disconnect(v.native(), cl)
}

// DisconnectKey is a wrapper around gtk_accel_group_disconnect_key().
func (v *AccelGroup) DisconnectKey(key uint, mods gdk.ModifierType) {
	C.gtk_accel_group_disconnect_key(v.native(), C.guint(key), C.GdkModifierType(mods))
}

// Lock is a wrapper around gtk_accel_group_lock().
func (v *AccelGroup) Lock() {
	C.gtk_accel_group_lock(v.native())
}

// Unlock is a wrapper around gtk_accel_group_unlock().
func (v *AccelGroup) Unlock() {
	C.gtk_accel_group_unlock(v.native())
}

// IsLocked is a wrapper around gtk_accel_group_get_is_locked().
func (v *AccelGroup) IsLocked() bool {
	return gobool(C.gtk_accel_group_get_is_locked(v.native()))
}

// AccelGroupFromClosure is a wrapper around gtk_accel_group_from_accel_closure().
func AccelGroupFromClosure(f interface{}) *AccelGroup {
	closure, _ := glib.ClosureNew(f)
	cl := (*C.GClosure)(unsafe.Pointer(closure))
	c := C.gtk_accel_group_from_accel_closure(cl)
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAccelGroup(obj)
}

// GetModifierMask is a wrapper around gtk_accel_group_get_modifier_mask().
func (v *AccelGroup) GetModifierMask() gdk.ModifierType {
	return gdk.ModifierType(C.gtk_accel_group_get_modifier_mask(v.native()))
}

// AccelGroupsActivate is a wrapper around gtk_accel_groups_activate().
func AccelGroupsActivate(obj *glib.Object, key uint, mods gdk.ModifierType) bool {
	o := C.toGObject(unsafe.Pointer(obj.Native()))
	return gobool(C.gtk_accel_groups_activate(o, C.guint(key), C.GdkModifierType(mods)))
}

// Activate is a wrapper around gtk_accel_group_activate().
func (v *AccelGroup) Activate(quark glib.Quark, acceleratable *glib.Object, key uint, mods gdk.ModifierType) bool {
	o := C.toGObject(unsafe.Pointer(acceleratable.Native()))
	return gobool(C.gtk_accel_group_activate(v.native(), C.GQuark(quark),
		o, C.guint(key), C.GdkModifierType(mods)))
}

// AccelGroupsFromObject is a wrapper around gtk_accel_groups_from_object().
func AccelGroupsFromObject(obj *glib.Object) *glib.SList {
	o := C.toGObject(unsafe.Pointer(obj.Native()))
	res := C.gtk_accel_groups_from_object(o)
	if res == nil {
		return nil
	}
	return (*glib.SList)(unsafe.Pointer(res))
}

/*
 * GtkAccelMap
 */

// AccelMap is a representation of GTK's GtkAccelMap.
type AccelMap struct {
	*glib.Object
}

// native returns a pointer to the underlying GtkAccelMap.
func (v *AccelMap) native() *C.GtkAccelMap {
	if v == nil {
		return nil
	}
	ptr := unsafe.Pointer(v.Object.Native())
	return C.toGtkAccelMap(ptr)
}

func marshalAccelMap(p uintptr) (interface{}, error) {
	c := C.g_value_get_object(C.toGValue(unsafe.Pointer(p)))
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAccelMap(obj), nil
}

func wrapAccelMap(obj *glib.Object) *AccelMap {
	return &AccelMap{obj}
}

// AccelMapAddEntry is a wrapper around gtk_accel_map_add_entry().
func AccelMapAddEntry(path string, key uint, mods gdk.ModifierType) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_add_entry((*C.gchar)(cstr), C.guint(key), C.GdkModifierType(mods))
}

type AccelKey struct {
	accelKey *C.GtkAccelKey
}

func (v *AccelKey) native() *C.GtkAccelKey {
	if v == nil {
		return nil
	}
	return v.accelKey
}

func wrapAccelKey(obj *C.GtkAccelKey) *AccelKey {
	return &AccelKey{obj}
}

// AccelMapLookupEntry is a wrapper around gtk_accel_map_lookup_entry().
func AccelMapLookupEntry(path string) *AccelKey {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	var v C.GtkAccelKey

	c := C.gtk_accel_map_lookup_entry((*C.gchar)(cstr), &v)
	if gobool(c) {
		return wrapAccelKey(&v)
	}
	return nil
}

// AccelMapChangeEntry is a wrapper around gtk_accel_map_change_entry().
func AccelMapChangeEntry(path string, key uint, mods gdk.ModifierType, replace bool) bool {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	return gobool(C.gtk_accel_map_change_entry((*C.gchar)(cstr), C.guint(key), C.GdkModifierType(mods), gbool(replace)))
}

// AccelMapLoad is a wrapper around gtk_accel_map_load().
func AccelMapLoad(fileName string) {
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_load((*C.gchar)(cstr))
}

// AccelMapSave is a wrapper around gtk_accel_map_save().
func AccelMapSave(fileName string) {
	cstr := C.CString(fileName)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_save((*C.gchar)(cstr))
}

// AccelMapLoadFD is a wrapper around gtk_accel_map_load_fd().
func AccelMapLoadFD(fd int) {
	C.gtk_accel_map_load_fd(C.gint(fd))
}

// AccelMapSaveFD is a wrapper around gtk_accel_map_save_fd().
func AccelMapSaveFD(fd int) {
	C.gtk_accel_map_save_fd(C.gint(fd))
}

// AccelMapAddFilter is a wrapper around gtk_accel_map_add_filter().
func AccelMapAddFilter(filter string) {
	cstr := C.CString(filter)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_add_filter((*C.gchar)(cstr))
}

// AccelMapGet is a wrapper around gtk_accel_map_get().
func AccelMapGet() *AccelMap {
	c := C.gtk_accel_map_get()
	if c == nil {
		return nil
	}
	obj := glib.Take(unsafe.Pointer(c))
	return wrapAccelMap(obj)
}

// AccelMapLockPath is a wrapper around gtk_accel_map_lock_path().
func AccelMapLockPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_lock_path((*C.gchar)(cstr))
}

// AccelMapUnlockPath is a wrapper around gtk_accel_map_unlock_path().
func AccelMapUnlockPath(path string) {
	cstr := C.CString(path)
	defer C.free(unsafe.Pointer(cstr))

	C.gtk_accel_map_unlock_path((*C.gchar)(cstr))
}

// These three functions are for system level access - thus not as high priority to implement
// TODO: void 	gtk_accelerator_parse_with_keycode ()
// TODO: gchar * 	gtk_accelerator_name_with_keycode ()
// TODO: gchar * 	gtk_accelerator_get_label_with_keycode ()

// TODO: GtkAccelKey * 	gtk_accel_group_find ()   - this function uses a function type - I don't know how to represent it in cgo
// TODO: gtk_accel_map_foreach_unfiltered  - can't be done without a function type
// TODO: gtk_accel_map_foreach  - can't be done without a function type

// TODO: gtk_accel_map_load_scanner
// TODO: gtk_widget_list_accel_closures

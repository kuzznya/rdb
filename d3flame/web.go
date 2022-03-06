package d3flame

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

var flameTmplData = &struct {
	D3Css        template.CSS
	D3Js         template.JS
	D3Flame      template.JS
	D3Tip        template.JS
	BootstrapCss template.CSS
}{
	D3Css:        template.CSS(d3Css),
	D3Js:         template.JS(d3Js),
	D3Flame:      template.JS(d3FlameGraphJs),
	D3Tip:        template.JS(d3TipJs),
	BootstrapCss: template.CSS(bootstrapCss),
}

func flamegraph(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("flamegraph").Parse(html))
	err := tmpl.Execute(w, flameTmplData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("500 - Internal Error"))
	}
}

// FlameItem is an Element in flamegraph
type FlameItem struct {
	Name     string   `json:"n"`
	Value    int      `json:"v"`
	Children children `json:"c,omitempty"`
}

type children map[string]*FlameItem

func (ch children) MarshalJSON() ([]byte, error) {
	list := make([]*FlameItem, 0, len(ch))
	for _, v := range ch {
		list = append(list, v)
	}
	return json.Marshal(list)
}

const testData = `
Production
{"c":[{"n":"genunixsyscall_mstate","v":89},{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"unixpage_lookup_create","v":1}],"n":"unixpage_lookup","v":1}],"n":"ufsufs_getpage","v":1}],"n":"genunixfop_getpage","v":1},{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"genunixpvn_plist_init","v":1},{"n":"unixlgrp_mem_choose","v":1},{"c":[{"c":[{"c":[{"n":"unixmutex_enter","v":1}],"n":"unixpage_get_mnode_freelist","v":1}],"n":"unixpage_get_freelist","v":1}],"n":"unixpage_create_va","v":1},{"c":[{"n":"unixpage_lookup_create","v":1}],"n":"unixpage_lookup","v":1}],"n":"genunixswap_getapage","v":4}],"n":"genunixswap_getpage","v":4}],"n":"genunixfop_getpage","v":4},{"c":[{"c":[{"n":"unixhwblkclr","v":3}],"n":"unixpfnzero","v":3}],"n":"unixpagezero","v":3}],"n":"genunixanon_zero","v":7}],"n":"genunixsegvn_faultpage","v":7},{"n":"ufsufs_getpage","v":1},{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"unixhment_compare","v":1}],"n":"genunixavl_find","v":1}],"n":"genunixavl_add","v":1}],"n":"unixhment_insert","v":2}],"n":"unixhment_assign","v":2}],"n":"unixhati_pte_map","v":2}],"n":"unixhati_load_common","v":2}],"n":"unixhat_memload","v":2}],"n":"unixhat_memload_region","v":2}],"n":"genunixsegvn_fault","v":11}],"n":"genunixas_fault","v":12},{"n":"genunixsegvn_fault","v":1}],"n":"unixpagefault","v":13}],"n":"unixtrap","v":13}],"n":"unix0xfffffffffb8001d6","v":13},{"n":"unix0xfffffffffb800c7c","v":42},{"n":"unix0xfffffffffb800c81","v":2},{"c":[{"n":"genunixgethrtime_unscaled","v":4},{"c":[{"c":[{"n":"unixtsc_gethrtimeunscaled","v":11},{"n":"unixtsc_read","v":186}],"n":"genunixgethrtime_unscaled","v":203},{"n":"unixtsc_gethrtimeunscaled","v":13}],"n":"genunixsyscall_mstate","v":355},{"n":"unixatomic_add_64","v":110}],"n":"unix0xfffffffffb800c86","v":472},{"c":[{"n":"genunixaudit_getstate","v":27},{"n":"genunixclear_stale_fd","v":10},{"n":"genunixdisp_lock_exit","v":27},{"c":[{"n":"FSSfss_preempt","v":1},{"n":"genunixaudit_getstate","v":15},{"n":"genunixclear_stale_fd","v":44},{"c":[{"n":"unixclear_int_flag","v":39},{"n":"unixdo_splx","v":1993},{"c":[{"c":[{"c":[{"n":"unixdo_splx","v":1}],"n":"genunixdisp_lock_exit_nopreempt","v":1}],"n":"unixpreempt","v":1}],"n":"unixkpreempt","v":1}],"n":"genunixdisp_lock_exit","v":2096},{"n":"genunixsigcheck","v":1},{"c":[{"n":"unixclear_int_flag","v":180},{"n":"unixsplr","v":400}],"n":"genunixthread_lock","v":670},{"n":"unixdo_splx","v":31},{"n":"unixi_ddi_splhigh","v":23},{"n":"unixlock_clear_splx","v":28},{"n":"unixlock_try","v":778},{"n":"unixlwp_getdatamodel","v":6},{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"unixtsc_gethrtimeunscaled","v":1}],"n":"genunixmstate_thread_onproc_time","v":1}],"n":"unixcaps_charge_adjust","v":1}],"n":"unixcpucaps_charge","v":3},{"c":[{"n":"unixcmt_balance","v":1},{"c":[{"n":"unixbitset_in_set","v":1}],"n":"unixcpu_wakeup_mwait","v":1}],"n":"unixsetbackdq","v":5}],"n":"FSSfss_preempt","v":8},{"n":"unixdo_splx","v":1},{"c":[{"n":"genunixdisp_lock_exit_high","v":1},{"c":[{"n":"unixmembar_enter","v":1}],"n":"unixdisp","v":1},{"n":"unixdo_splx","v":1},{"c":[{"c":[{"n":"genunixschedctl_save","v":1}],"n":"genunixsavectx","v":2}],"n":"unixresume","v":2}],"n":"unixswtch","v":5}],"n":"unixpreempt","v":14},{"n":"unixprunstop","v":36},{"n":"unixsplr","v":92},{"n":"unixsplx","v":6}],"n":"genunixpost_syscall","v":4245},{"n":"genunixthread_lock","v":33},{"n":"unixlwp_getdatamodel","v":3},{"n":"unixprunstop","v":2}],"n":"unix0xfffffffffb800c91","v":4361},{"c":[{"n":"genunixgethrtime_unscaled","v":7},{"c":[{"c":[{"n":"unixtsc_gethrtimeunscaled","v":17},{"n":"unixtsc_read","v":160}],"n":"genunixgethrtime_unscaled","v":182},{"n":"unixtsc_gethrtimeunscaled","v":12}],"n":"genunixsyscall_mstate","v":412},{"n":"unixatomic_add_64","v":95}],"n":"unix0xfffffffffb800ca0","v":517},{"n":"unix_sys_rtt","v":6},{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"genunixcpu_decay","v":1}],"n":"genunixcpu_grow","v":1}],"n":"genunixcpu_update_pct","v":1}],"n":"genunixnew_mstate","v":1}],"n":"unixtrap","v":1}],"n":"unixsys_rtt_common","v":1}],"n":"unix_sys_rtt_ints_disabled","v":1},{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"doorfsdoor_close","v":1}],"n":"nfsnm_close","v":1}],"n":"genunixfop_close","v":1}],"n":"genunixclosef","v":1}],"n":"genunixclose_exec","v":1}],"n":"genunixexec_common","v":1}],"n":"genunixexece","v":1}],"n":"unix_sys_sysenter_post_swapgs","v":1},{"c":[{"n":"genunixgethrtime_unscaled","v":11},{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"unixmtype_func","v":1},{"n":"unixmutex_enter","v":1}],"n":"unixpage_get_mnode_freelist","v":2}],"n":"unixpage_get_freelist","v":2}],"n":"unixpage_create_va","v":3}],"n":"genunixpvn_read_kluster","v":3}],"n":"ufsufs_getpage_ra","v":3}],"n":"ufsufs_getpage","v":3}],"n":"genunixfop_getpage","v":3}],"n":"genunixsegvn_faulta","v":3}],"n":"genunixas_faulta","v":3}],"n":"genunixmemcntl","v":3},{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"unixhtable_lookup","v":1}],"n":"unixhtable_walk","v":1}],"n":"unixhat_unload_callback","v":1}],"n":"genunixsegvn_unmap","v":1}],"n":"genunixas_unmap","v":1}],"n":"unixmmapobj_map_elf","v":1}],"n":"unixmmapobj_map_interpret","v":1}],"n":"unixmmapobj","v":1}],"n":"genunixmmapobjsys","v":1},{"c":[{"n":"genunixcopen","v":7},{"c":[{"n":"genunixaudit_getstate","v":62},{"c":[{"n":"genunixaudit_falloc","v":8},{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"unixswtch","v":1}],"n":"unixpreempt","v":1}],"n":"unixkpreempt","v":1}],"n":"unixsys_rtt_common","v":1}],"n":"unix_sys_rtt_ints_disabled","v":1}],"n":"genunixaudit_getstate","v":66},{"n":"genunixaudit_unfalloc","v":32},{"n":"genunixcrfree","v":9},{"n":"genunixcrhold","v":5},{"n":"genunixcv_broadcast","v":16},{"c":[{"c":[{"n":"genunixkmem_cache_alloc","v":11},{"c":[{"n":"genunixkmem_cache_alloc","v":66},{"n":"unixmutex_enter","v":122},{"n":"unixmutex_exit","v":46}],"n":"genunixkmem_zalloc","v":280},{"n":"unixbzero","v":8}],"n":"genunixaudit_falloc","v":313},{"n":"genunixcrhold","v":11},{"n":"genunixkmem_cache_alloc","v":49},{"n":"genunixkmem_zalloc","v":13},{"c":[{"n":"genunixfd_find","v":13},{"n":"genunixfd_reserve","v":9},{"c":[{"n":"genunixfd_find","v":161},{"n":"genunixfd_reserve","v":15}],"n":"genunixufalloc_file","v":294},{"n":"unixmutex_enter","v":197},{"n":"unixmutex_exit","v":29}],"n":"genunixufalloc","v":551},{"n":"genunixufalloc_file","v":20},{"n":"unixatomic_add_32","v":134},{"n":"unixmutex_enter","v":99},{"n":"unixmutex_exit","v":58}],"n":"genunixfalloc","v":1363},{"n":"genunixfd_reserve","v":8},{"n":"genunixkmem_cache_alloc","v":9},{"n":"genunixkmem_cache_free","v":5},{"n":"genunixlookupnat","v":69},{"n":"genunixset_errno","v":24},{"c":[{"n":"genunixaudit_getstate","v":31},{"n":"genunixcv_broadcast","v":25},{"n":"genunixfd_reserve","v":35}],"n":"genunixsetf","v":187},{"n":"genunixufalloc","v":10},{"c":[{"c":[{"n":"genunixkmem_cache_free","v":5},{"c":[{"n":"genunixkmem_cache_free","v":73},{"n":"unixmutex_enter","v":111},{"n":"unixmutex_exit","v":55}],"n":"genunixkmem_free","v":288}],"n":"genunixaudit_unfalloc","v":340},{"n":"genunixcrfree","v":13},{"n":"genunixkmem_cache_free","v":51},{"n":"genunixkmem_free","v":11},{"n":"unixatomic_add_32_nv","v":100},{"n":"unixmutex_enter","v":97},{"n":"unixmutex_exit","v":56}],"n":"genunixunfalloc","v":729},{"c":[{"c":[{"c":[{"c":[{"n":"genunixaudit_getstate","v":16},{"n":"genunixfop_lookup","v":55},{"c":[{"n":"genunixaudit_getstate","v":21},{"n":"genunixcrgetmapped","v":55},{"n":"genunixfop_inactive","v":39},{"c":[{"n":"genunixcrgetmapped","v":57},{"n":"genunixdnlc_lookup","v":26},{"n":"genunixfop_lookup","v":85},{"n":"genunixkmem_alloc","v":73},{"n":"genunixtraverse","v":30},{"n":"genunixvfs_matchops","v":28},{"c":[{"c":[{"n":"genunixkmem_cache_alloc","v":241},{"n":"unixmutex_enter","v":366},{"n":"unixmutex_exit","v":149}],"n":"genunixkmem_alloc","v":934},{"n":"genunixkmem_cache_alloc","v":32}],"n":"genunixvn_setpath","v":1969},{"c":[{"n":"genunixcrgetmapped","v":36},{"c":[{"n":"genunixcrgetmapped","v":58},{"n":"genunixdnlc_lookup","v":70},{"n":"genunixvn_rele","v":14},{"n":"ufsufs_iaccess","v":91},{"c":[{"n":"genunixcrgetuid","v":30},{"c":[{"n":"genunixmemcmp","v":38},{"c":[{"n":"genunixmemcmp","v":277}],"n":"unixbcmp","v":295}],"n":"genunixdnlc_lookup","v":1843},{"n":"genunixsecpolicy_vnode_access2","v":72},{"n":"genunixvn_rele","v":39},{"c":[{"n":"genunixcrgetuid","v":22},{"n":"genunixsecpolicy_vnode_access2","v":217}],"n":"ufsufs_iaccess","v":648},{"n":"unixbcmp","v":42},{"n":"unixmutex_enter","v":980},{"n":"unixmutex_exit","v":350},{"n":"unixrw_enter","v":525},{"n":"unixrw_exit","v":439}],"n":"ufsufs_lookup","v":5399}],"n":"genunixfop_lookup","v":6470},{"n":"genunixkmem_cache_alloc","v":39},{"c":[{"n":"genunixrwst_exit","v":18},{"n":"genunixrwst_tryenter","v":32},{"n":"genunixvn_mountedvfs","v":11},{"n":"genunixvn_vfslocks_getlock","v":62},{"n":"genunixvn_vfslocks_rele","v":50},{"c":[{"n":"genunixkmem_alloc","v":32},{"n":"genunixrwst_enter_common","v":32},{"n":"genunixrwst_init","v":28},{"c":[{"n":"genunixrwst_enter_common","v":264},{"n":"unixmutex_enter","v":337},{"n":"unixmutex_exit","v":105}],"n":"genunixrwst_tryenter","v":734},{"c":[{"n":"genunixcv_init","v":53},{"c":[{"c":[{"n":"genunixkmem_cpu_reload","v":2}],"n":"genunixkmem_cache_alloc","v":168},{"n":"unixmutex_enter","v":379},{"n":"unixmutex_exit","v":155}],"n":"genunixkmem_alloc","v":795},{"n":"genunixkmem_cache_alloc","v":29},{"c":[{"n":"genunixcv_init","v":65},{"n":"unixmutex_init","v":53}],"n":"genunixrwst_init","v":236},{"n":"unixmutex_init","v":46}],"n":"genunixvn_vfslocks_getlock","v":1357},{"n":"unixmutex_enter","v":727},{"n":"unixmutex_exit","v":371}],"n":"genunixvn_vfsrlock","v":3342},{"c":[{"n":"genunixcv_broadcast","v":25},{"n":"genunixkmem_free","v":35},{"n":"genunixrwst_destroy","v":32},{"c":[{"n":"genunixcv_broadcast","v":40}],"n":"genunixrwst_exit","v":167},{"n":"genunixvn_vfslocks_getlock","v":120},{"c":[{"n":"genunixcv_destroy","v":77},{"n":"genunixkmem_cache_free","v":22},{"c":[{"n":"genunixkmem_cache_free","v":154},{"n":"unixmutex_enter","v":316},{"n":"unixmutex_exit","v":148}],"n":"genunixkmem_free","v":693},{"c":[{"n":"genunixcv_destroy","v":42},{"n":"unixmutex_destroy","v":176}],"n":"genunixrwst_destroy","v":296},{"n":"unixmutex_destroy","v":31}],"n":"genunixvn_vfslocks_rele","v":1420},{"n":"unixmutex_enter","v":1202},{"n":"unixmutex_exit","v":512}],"n":"genunixvn_vfsunlock","v":3578}],"n":"genunixtraverse","v":7243},{"n":"genunixvfs_getops","v":21},{"c":[{"n":"genunixvfs_getops","v":157},{"n":"unixmembar_consumer","v":123}],"n":"genunixvfs_matchops","v":336},{"n":"genunixvn_alloc","v":20},{"n":"genunixvn_exists","v":17},{"n":"genunixvn_mountedvfs","v":30},{"n":"genunixvn_setops","v":41},{"n":"genunixvn_vfsrlock","v":13},{"n":"genunixvn_vfsunlock","v":40},{"n":"lofslfind","v":26},{"n":"lofslsave","v":27},{"n":"lofsmakelfsnode","v":28},{"c":[{"n":"genunixkmem_cache_alloc","v":234},{"n":"genunixkmem_cpu_reload","v":1},{"c":[{"n":"genunixkmem_cache_alloc","v":179},{"n":"genunixvn_recycle","v":33},{"c":[{"c":[{"n":"genunixvsd_free","v":155}],"n":"genunixvn_recycle","v":319},{"n":"genunixvsd_free","v":14}],"n":"genunixvn_reinit","v":424},{"n":"unixmutex_enter","v":318},{"n":"unixmutex_exit","v":142}],"n":"genunixvn_alloc","v":1189},{"n":"genunixvn_exists","v":50},{"n":"genunixvn_reinit","v":48},{"n":"genunixvn_setops","v":160},{"n":"lofslfind","v":278},{"n":"lofslsave","v":162},{"n":"lofsmakelfsnode","v":82},{"n":"lofstable_lock_enter","v":220},{"n":"unixatomic_cas_64","v":318},{"n":"unixmembar_consumer","v":237},{"n":"unixmutex_enter","v":640},{"n":"unixmutex_exit","v":138}],"n":"lofsmakelonode","v":4212},{"n":"lofstable_lock_enter","v":43},{"n":"ufsufs_lookup","v":46},{"n":"unixatomic_add_32","v":325},{"n":"unixmutex_exit","v":26}],"n":"lofslo_lookup","v":19887},{"n":"lofsmakelonode","v":39},{"n":"unixbcopy","v":896},{"n":"unixmutex_enter","v":947},{"n":"unixmutex_exit","v":337},{"c":[{"c":[{"c":[{"n":"unixdispatch_hilevel","v":1}],"n":"unixdo_interrupt","v":1}],"n":"unix_interrupt","v":1}],"n":"unixstrlen","v":2659},{"n":"zfsspecvp_check","v":10},{"n":"zfszfs_fastaccesschk_execute","v":4},{"c":[{"n":"genunixcrgetuid","v":6},{"c":[{"n":"genunixmemcmp","v":3},{"c":[{"n":"genunixmemcmp","v":38}],"n":"unixbcmp","v":45}],"n":"genunixdnlc_lookup","v":263},{"n":"unixbcmp","v":11},{"n":"unixmutex_enter","v":309},{"n":"unixmutex_exit","v":135},{"n":"zfsspecvp_check","v":20},{"c":[{"n":"genunixcrgetuid","v":2}],"n":"zfszfs_fastaccesschk_execute","v":50}],"n":"zfszfs_lookup","v":946}],"n":"genunixfop_lookup","v":29216},{"n":"genunixfsop_root","v":62},{"n":"genunixpn_fixslash","v":44},{"n":"genunixpn_getcomponent","v":454},{"c":[{"c":[{"n":"lofslo_root","v":80},{"n":"unixmutex_enter","v":95},{"n":"unixmutex_exit","v":59}],"n":"genunixfsop_root","v":297},{"n":"genunixrwst_exit","v":12},{"n":"genunixrwst_tryenter","v":37},{"n":"genunixvn_mountedvfs","v":20},{"n":"genunixvn_rele","v":19},{"n":"genunixvn_vfslocks_getlock","v":47},{"n":"genunixvn_vfslocks_rele","v":34},{"c":[{"n":"genunixkmem_alloc","v":11},{"n":"genunixrwst_enter_common","v":28},{"n":"genunixrwst_init","v":13},{"c":[{"n":"genunixrwst_enter_common","v":314},{"n":"unixmutex_enter","v":238},{"n":"unixmutex_exit","v":49}],"n":"genunixrwst_tryenter","v":628},{"c":[{"n":"genunixcv_init","v":56},{"c":[{"n":"genunixkmem_cache_alloc","v":126},{"n":"unixmutex_enter","v":252},{"n":"unixmutex_exit","v":95}],"n":"genunixkmem_alloc","v":533},{"n":"genunixkmem_cache_alloc","v":17},{"c":[{"n":"genunixcv_init","v":49},{"n":"unixmutex_init","v":38}],"n":"genunixrwst_init","v":173},{"n":"unixmutex_init","v":31}],"n":"genunixvn_vfslocks_getlock","v":973},{"n":"unixmutex_enter","v":455},{"n":"unixmutex_exit","v":250}],"n":"genunixvn_vfsrlock","v":2414},{"c":[{"n":"genunixcv_broadcast","v":14},{"n":"genunixkmem_free","v":17},{"n":"genunixrwst_destroy","v":20},{"c":[{"n":"genunixcv_broadcast","v":19}],"n":"genunixrwst_exit","v":110},{"n":"genunixvn_vfslocks_getlock","v":79},{"c":[{"n":"genunixcv_destroy","v":81},{"n":"genunixkmem_cache_free","v":18},{"c":[{"n":"genunixkmem_cache_free","v":116},{"n":"unixmutex_enter","v":195},{"n":"unixmutex_exit","v":90}],"n":"genunixkmem_free","v":457},{"c":[{"n":"genunixcv_destroy","v":31},{"n":"unixmutex_destroy","v":53}],"n":"genunixrwst_destroy","v":146},{"n":"unixmutex_destroy","v":17}],"n":"genunixvn_vfslocks_rele","v":903},{"n":"unixmutex_enter","v":823},{"n":"unixmutex_exit","v":356}],"n":"genunixvn_vfsunlock","v":2372},{"n":"lofslo_root","v":31},{"n":"unixmutex_enter","v":95},{"n":"unixmutex_exit","v":56}],"n":"genunixtraverse","v":5557},{"n":"genunixvn_mountedvfs","v":43},{"c":[{"n":"genunixcrgetmapped","v":31},{"c":[{"n":"genunixcrgetmapped","v":41},{"n":"lofsfreelonode","v":35},{"c":[{"n":"genunixkmem_cache_free","v":29},{"n":"genunixvn_free","v":26},{"n":"genunixvn_invalid","v":20},{"n":"genunixvn_rele","v":25},{"c":[{"c":[{"n":"genunixkmem_cpu_reload","v":1}],"n":"genunixkmem_cache_free","v":184},{"n":"genunixkmem_free","v":115},{"c":[{"c":[{"n":"genunixkmem_cpu_reload","v":4}],"n":"genunixkmem_cache_free","v":215},{"n":"genunixkmem_cpu_reload","v":5},{"c":[{"n":"genunixkmem_cache_free","v":209},{"n":"unixmutex_enter","v":299},{"n":"unixmutex_exit","v":160}],"n":"genunixkmem_free","v":785},{"n":"genunixvsd_free","v":48},{"n":"unixmutex_enter","v":314},{"n":"unixmutex_exit","v":171}],"n":"genunixvn_free","v":1663},{"n":"genunixvn_invalid","v":47},{"n":"genunixvn_rele","v":64},{"n":"genunixvsd_free","v":17},{"n":"lofstable_lock_enter","v":189},{"n":"unixmembar_consumer","v":106},{"n":"unixmutex_enter","v":905},{"n":"unixmutex_exit","v":358},{"n":"unixstrlen","v":1238}],"n":"lofsfreelonode","v":5313},{"n":"lofstable_lock_enter","v":44},{"n":"unixatomic_add_32","v":292},{"n":"unixmutex_enter","v":279},{"n":"unixmutex_exit","v":212}],"n":"lofslo_inactive","v":6307}],"n":"genunixfop_inactive","v":6689},{"n":"lofslo_inactive","v":21}],"n":"genunixvn_rele","v":6943},{"n":"genunixvn_setpath","v":58},{"n":"genunixvn_vfsrlock","v":12},{"n":"genunixvn_vfsunlock","v":20},{"n":"lofslo_lookup","v":65},{"n":"unixmutex_enter","v":575},{"n":"unixmutex_exit","v":379},{"n":"unixstrlen","v":107},{"n":"zfszfs_lookup","v":22}],"n":"genunixlookuppnvp","v":44242},{"n":"genunixpn_fixslash","v":14},{"n":"genunixpn_getcomponent","v":41},{"n":"genunixtraverse","v":17},{"n":"genunixvn_mountedvfs","v":56},{"n":"genunixvn_rele","v":73},{"c":[{"n":"unixmutex_delay_default","v":1},{"n":"unixtsc_read","v":1}],"n":"unixmutex_vector_enter","v":2}],"n":"genunixlookuppnatcred","v":44681},{"n":"genunixlookuppnvp","v":10},{"c":[{"n":"unixcopyinstr","v":25},{"n":"unixcopystr","v":598}],"n":"genunixpn_get_buf","v":687},{"n":"unixcopyinstr","v":18},{"n":"unixmutex_enter","v":320},{"n":"unixmutex_exit","v":163}],"n":"genunixlookupnatcred","v":45978},{"n":"genunixlookuppnatcred","v":12},{"n":"genunixpn_get_buf","v":13}],"n":"genunixlookupnat","v":46075},{"n":"genunixlookupnatcred","v":22}],"n":"genunixvn_openat","v":46342},{"n":"unixmutex_enter","v":303},{"n":"unixmutex_exit","v":38}],"n":"genunixcopen","v":49444},{"n":"genunixfalloc","v":36},{"n":"genunixset_errno","v":9},{"n":"genunixsetf","v":16},{"n":"genunixunfalloc","v":39},{"n":"genunixvn_openat","v":14}],"n":"genunixopenat","v":49647}],"n":"genunixopen","v":49669},{"n":"genunixopenat","v":17},{"c":[{"c":[{"c":[{"n":"genunixdotoprocs","v":1}],"n":"genunixdoprio","v":1}],"n":"genunixpriocntl_common","v":1}],"n":"genunixpriocntlsys","v":1},{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"genunixdnlc_lookup","v":1}],"n":"ufsufs_lookup","v":1}],"n":"genunixfop_lookup","v":1}],"n":"lofslo_lookup","v":1}],"n":"genunixfop_lookup","v":1}],"n":"genunixlookuppnvp","v":1}],"n":"genunixlookuppnatcred","v":1}],"n":"genunixlookuppn","v":1}],"n":"genunixresolvepath","v":1},{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"genunixkmem_cache_free","v":1}],"n":"genunixkmem_free","v":1}],"n":"genunixremovectx","v":1}],"n":"genunixschedctl_lwp_cleanup","v":1}],"n":"genunixexitlwps","v":1},{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"c":[{"n":"unixhment_compare","v":2}],"n":"genunixavl_find","v":2}],"n":"unixhment_remove","v":2},{"n":"unixpage_numtopp_nolock","v":1}],"n":"unixhat_pte_unmap","v":3}],"n":"unixhat_unload_callback","v":3}],"n":"genunixsegvn_unmap","v":3}],"n":"genunixas_free","v":3}],"n":"genunixrelvm","v":3},{"c":[{"c":[{"c":[{"c":[{"n":"genunixvmem_free","v":1}],"n":"genunixsegkp_release_internal","v":1}],"n":"genunixsegkp_release","v":1}],"n":"genunixschedctl_freepage","v":1}],"n":"genunixschedctl_proc_cleanup","v":1}],"n":"genunixproc_exit","v":5}],"n":"genunixexit","v":5}],"n":"genunixrexit","v":5},{"c":[{"c":[{"n":"unixtsc_gethrtimeunscaled","v":43},{"n":"unixtsc_read","v":367}],"n":"genunixgethrtime_unscaled","v":420},{"n":"unixtsc_gethrtimeunscaled","v":59}],"n":"genunixsyscall_mstate","v":1336},{"n":"unixatomic_add_64","v":205}],"n":"unixsys_syscall","v":51908}],"n":"root","v":57412}
`

// Web starts a web server to render flamegraph
func Web(data []byte, port int) chan<- struct{} {
	server := &http.Server{
		Addr: ":" + strconv.Itoa(port),
	}
	http.HandleFunc("/flamegraph", flamegraph)
	http.HandleFunc("/stacks.json", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(data)
	})
	stop := make(chan struct{})
	go func() {
		<-stop
		_ = server.Close()
	}()
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
	return stop
}

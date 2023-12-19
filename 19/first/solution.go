package main

import (
	"fmt"
	"strconv"
	"strings"
	"syscall"
	"time"
)

func main() {

	startTime := time.Now()

	var rows = strings.Split(workflows, "\n")
	mapWrkflws := make(map[string]Workflow)
	for _, row := range rows {
		wrkflw := createWorkflow(row)
		mapWrkflws[wrkflw.name] = wrkflw
	}

	rows = strings.Split(input, "\n")
	res := 0

	for _, row := range rows {
		part := createPart(row)
		wrkflw := mapWrkflws["in"]
		newWrkflw := ""
		for {
			newWrkflw = calcWrk(wrkflw, part)
			if newWrkflw == "A" || newWrkflw == "R" {
				break
			}
			wrkflw2, ok := mapWrkflws[newWrkflw]
			if !ok {
				fmt.Errorf("Not exists workflow", newWrkflw)
				syscall.Exit(1)
			}
			wrkflw = wrkflw2
		}
		if newWrkflw == "A" {
			res += part.x + part.m + part.a + part.s
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Your function took %s\n", elapsedTime)

	fmt.Println(res)
}

func calcWrk(wrkflw Workflow, part Part) string {
	for _, operation := range wrkflw.operations {
		value := 0
		if operation.isDefault {
			return operation.reference
		} else if operation.field == "x" {
			value = part.x
		} else if operation.field == "m" {
			value = part.m
		} else if operation.field == "a" {
			value = part.a
		} else {
			value = part.s
		}

		if operation.symbol == ">" && value > operation.value {
			return operation.reference
		} else if operation.symbol == "<" && value < operation.value {
			return operation.reference
		}
	}
	return ""
}

func createPart(row string) Part {
	row = strings.Replace(row, "{", "", 1)
	row = strings.Replace(row, "}", "", 1)

	splitted := strings.Split(row, ",")
	var part Part
	for _, el := range splitted {
		ell := strings.Split(el, "=")
		value, err := strconv.Atoi(ell[1])
		if err != nil {
			fmt.Errorf("ERROR", err)
			syscall.Exit(1)
		}
		if ell[0] == "x" {
			part.x = value
		} else if ell[0] == "m" {
			part.m = value
		} else if ell[0] == "a" {
			part.a = value
		} else {
			part.s = value
		}
	}
	return part
}

func createWorkflow(row string) Workflow {
	parts := strings.Split(row, "{")
	name := parts[0]
	parts2 := strings.Replace(parts[1], "}", "", 1)
	parts3 := strings.Split(parts2, ",")

	var operations []Operation
	for _, part := range parts3 {
		operations = append(operations, createOperation(part))
	}

	return Workflow{name, operations}
}

func createOperation(input string) Operation {
	if !strings.Contains(input, ":") {
		return Operation{true, input, "", "", 0}
	}
	parts := strings.Split(input, ":")
	reference := parts[1]
	field := string(parts[0][0])
	symbol := string(parts[0][1])
	valueStr := parts[0][2:]
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		fmt.Errorf("ERROR", err)
		syscall.Exit(1)
	}
	return Operation{false, reference, field, symbol, value}
}

type Workflow struct {
	name       string
	operations []Operation
}

type Operation struct {
	isDefault bool
	reference string
	field     string
	symbol    string
	value     int
}

type Part struct {
	x int
	m int
	a int
	s int
}

const workflows = `sp{x<3224:A,a>1743:A,m<3854:A,R}
qzq{m>2723:R,xm}
gzj{x>2762:A,a<541:br,R}
chj{s>2340:hx,x>2551:vbb,s>1050:gf,jdz}
jhf{m<2021:ht,s<2175:zjp,x<3626:nkb,mzl}
lb{a>3686:A,s<1789:R,R}
zpg{m>2216:R,m>1278:R,R}
ptv{m<289:A,a<1547:A,x>2091:A,A}
qsx{m>3295:A,x<651:A,x>1078:gj,vsr}
bcx{a<597:R,x<3112:A,s>2486:R,A}
rbm{s<3766:R,A}
jx{m>2597:A,m<1086:A,a<2643:A,A}
jxq{s<487:R,s<689:R,R}
kk{a>3590:txm,x>3449:kzx,dxf}
vrt{x<2361:nfd,s>2509:zbx,s>1548:nk,vvx}
kcd{a<1577:R,x>502:R,R}
fjm{m>2560:R,m>1928:R,a<2849:R,A}
msb{s>2580:R,x>3839:A,A}
zhs{m>326:A,s>2498:A,x>163:A,R}
zj{m<231:R,a>1349:A,R}
zjp{x<3517:rfz,s<1052:rgd,x>3741:bqr,nrd}
ghp{x<1955:A,x>2826:R,s>3066:A,A}
ksm{m>1502:R,m>882:A,a<606:A,A}
zsg{x>3633:A,A}
fdr{x<2575:A,R}
xg{m>2572:A,a<59:R,A}
vq{s<411:R,a<2945:A,R}
gkr{s>3186:A,x>3478:nfb,s>2590:dv,R}
jj{m>351:R,x<2330:R,A}
fr{s>1144:R,s>1127:A,R}
rt{a>3344:R,R}
ln{m>3309:A,A}
ch{s<3179:R,a>2826:R,rd}
jmb{a>3549:R,R}
lhp{s<2235:A,m>2701:R,m>2626:A,R}
tt{a<3326:R,a>3698:A,A}
sl{x<782:nbh,s>994:lkq,hnx}
xsp{m<832:A,A}
bcb{m<1059:R,A}
jzk{a<392:A,A}
qr{x>1414:A,s<2146:A,A}
lzl{a<394:A,R}
df{s>590:R,s<332:xzt,R}
vls{m<1602:qlt,m<2715:lht,s<2608:tvf,fdd}
xxp{s<3569:nqg,a<777:R,x<756:mjr,A}
qpx{x<1460:A,A}
qn{s>1487:A,m<675:R,x>3304:R,A}
lqf{s>2544:R,m<2472:gz,R}
bs{m>3019:R,m>2339:R,R}
tkg{s<2861:R,R}
sr{a>534:R,a>185:rrq,A}
pr{x<2906:tl,A}
zk{a<2476:qgr,s<2925:fns,s<3101:xkg,ch}
kf{x<3638:A,R}
xcr{m>2307:R,x<3357:R,a<2010:A,A}
rkq{a<688:mjx,a<959:A,R}
bh{m<947:zr,s>1275:fl,R}
lm{m>2643:qqn,a>3231:A,pkf}
tr{s>2521:R,R}
lfs{s>769:fkb,s>712:A,A}
xz{a>3913:A,a<3873:R,A}
mjx{s>1084:A,x>2705:A,s<923:R,R}
xlb{x>3614:dqp,m>2061:A,cdv}
zqq{s>1046:hm,sr}
xfr{m<3014:qnv,fgl}
dbt{a<956:R,s<2691:R,s<2752:R,R}
fhh{x<1734:txc,a>2962:mdk,cg}
gq{a>146:A,s<3234:R,x<3295:A,R}
qf{x<3200:R,x<3230:R,s>3278:R,A}
cdt{x>3462:gzv,a<562:R,x>3217:gg,R}
xzt{a<2851:R,x>1470:A,m>3329:R,A}
sj{a<1691:A,x<3862:A,a<1866:A,R}
hsm{x<1953:A,A}
xqb{a>2533:A,R}
qb{x>778:gpg,s>3739:fn,kcd}
xzr{x>2124:A,R}
tnx{a<3540:R,m>951:R,R}
xmh{m<1107:R,a>639:R,R}
gf{a>3197:vg,pmt}
kbz{m<2517:js,x>2496:pcn,x>2425:hf,dgj}
qkp{m<3678:A,R}
trb{a>1335:A,m<2394:A,A}
rcm{a<2390:xgz,s<1270:tfn,m<2577:xlb,xld}
kq{x<2984:A,m<601:A,R}
bc{a>646:R,A}
sjl{x>3105:R,A}
gpg{m<3530:qln,qrc}
lxs{x>3313:R,s<3288:A,kv}
mv{x>1402:A,s<2939:A,m>959:xqb,tz}
lfd{m>2983:R,s<2543:A,m<2873:R,R}
tgn{s>3028:A,s<2954:A,m>3355:A,R}
tf{x>869:fhh,x<448:vs,x>684:hsr,xv}
qnk{s<2742:xhz,mv}
qt{x>3036:vm,a>3434:R,A}
mgl{x>431:A,m>2631:R,A}
bzb{s<3312:A,s<3720:hb,cb}
rbt{x>1240:A,x<1041:A,s<2792:A,A}
lj{x<3081:A,x>3573:fv,R}
cp{a<3760:A,x<104:R,a<3871:A,R}
sn{a<610:zpg,A}
sd{a<3746:A,R}
zmg{m<1681:R,s<3671:R,R}
ghl{s>1536:R,x<3612:A,s>1287:R,R}
mq{a<563:A,s<347:R,x<149:R,A}
dsx{s>841:zls,nx}
vs{a<3376:jk,hcs}
fv{a<3819:A,a>3888:R,s<2966:R,A}
rd{s<3246:R,R}
nk{x<2560:nrz,ssz}
mdk{x<2220:R,R}
pf{s<2767:R,R}
pmv{m>2273:A,m>1913:R,A}
mc{m>402:qj,s<3604:R,x<507:A,A}
zkp{a<1671:A,A}
jdz{m<1778:tf,ttt}
lht{s>2762:sx,m<2034:R,fm}
jp{m<2417:R,x>3602:R,A}
hg{x<3816:A,a<3754:R,m<636:R,A}
ts{x<1201:A,s<299:A,R}
hdg{a<884:A,R}
sth{x>898:R,x>517:A,a<447:R,A}
vc{x>833:A,R}
ds{a<1515:R,R}
pmt{s<1724:jm,m>2587:cl,qrz}
qc{x<733:A,R}
ssz{s>1944:hsp,A}
td{m<2680:A,A}
ql{m>1880:lpl,s>1345:R,pm}
kh{x<1677:rrd,s<777:vlz,A}
njl{m>2017:A,a<624:A,R}
xxx{x<3667:rc,s>1145:hg,a<3727:R,rm}
nxs{m>3620:R,m<3459:R,A}
gj{m>3010:R,s<3438:R,x>1282:A,R}
xvg{a<3524:A,s<418:A,m<2681:A,R}
mjr{s>3683:A,a<1639:A,x>341:A,A}
mbr{m>3538:R,x<897:R,s>2681:R,R}
mmg{a<3144:A,A}
kn{m>1078:A,s<1887:R,R}
ht{m<936:nl,x<3389:vjj,cd}
tn{m<1847:tj,s<3696:nbt,s>3815:svc,rbm}
ncf{a<332:A,xmh}
lc{x<293:A,s<1777:A,a>3569:R,R}
gqt{s<638:R,A}
gtk{x<1867:lb,A}
txf{s>1100:fr,R}
hkn{s>3748:czr,m<2175:xxp,zx}
zpx{m<3448:A,a<3506:A,A}
dlj{x>3324:A,m>2976:A,R}
vm{a<3394:R,a<3772:R,R}
gcp{m<3549:R,s<3753:A,m<3801:A,sth}
lck{x<1530:rcj,th}
vp{m<2446:A,pb}
rxs{m>2148:xfr,x>2413:bnd,m<925:mn,bpd}
rr{m>2769:A,A}
lp{a<3631:dd,A}
jl{m<2480:A,R}
ttp{a<1440:A,R}
bvr{m<1376:R,a<2606:A,x>312:A,A}
nfd{s>2564:rh,s<1039:mcj,m<2531:qd,bf}
cd{x<3668:R,x>3840:ttp,x>3765:cf,A}
ctk{x<362:R,a>1685:A,x>543:R,A}
lx{x>944:rmb,a<856:ncf,x<341:rgt,jkc}
xn{a<528:R,s<698:R,x>2513:A,R}
ld{a>3069:R,a>2737:R,A}
lf{x>636:R,m>1627:bs,A}
btr{m>2761:czv,s<2146:A,lhp}
sc{s<3669:A,x>3260:A,a>1921:R,R}
fgl{a<1538:jkd,x>2095:R,m>3384:A,zq}
tvf{a>646:R,lq}
dxf{s>1265:pgb,s>591:jdh,lm}
pg{a>2684:kcs,ln}
nl{s<1687:A,s<2645:tfk,zkp}
mhx{x>3495:R,A}
djk{m>3229:R,R}
vpf{s>467:tnx,A}
hxc{a<1252:R,a<1437:R,A}
hn{x<1490:rz,x<2950:dtk,lr}
jm{s>1371:vp,s<1194:txf,a>2796:rg,hl}
zqk{s<732:A,a>3079:R,R}
hj{x<3237:R,A}
mlj{a<1660:R,A}
hsp{a>505:R,a>190:A,a>109:R,R}
fng{s<411:qc,dbq}
htk{s>2864:pzb,s<2590:lx,x>908:kp,vdr}
ck{x>1172:A,a>809:R,R}
dxd{s>2128:R,R}
vvx{s>798:rkq,s<425:sn,x<2568:qmb,ff}
bn{x>3635:R,A}
qgr{m<3073:A,R}
lxp{x>789:R,a>968:cgj,a<346:gm,R}
lhd{a>1548:phc,a<1397:gc,s>649:A,ls}
qhj{m<3630:A,s<2526:A,x<3209:A,R}
zz{a<832:R,a>852:A,A}
dm{m>685:A,x<398:A,m<325:R,A}
xtb{m<59:R,R}
mxt{a<959:A,s>448:A,A}
zf{s<1534:A,a<2456:A,m>870:R,R}
kcs{a<2970:R,A}
nbt{s<3490:A,s<3621:R,jmb}
xf{m>1414:R,a>758:R,A}
trh{s>1864:fjm,A}
gg{m<1992:A,A}
lq{a>634:A,R}
mmb{a>956:A,A}
nbh{s>1081:jmk,a>1242:jkl,s>435:tnz,rb}
rcj{m<2679:R,a>3271:tkg,mbr}
rl{s>2985:R,A}
tnz{x<505:R,A}
kv{m>2334:R,x>3309:R,R}
in{a>2167:chj,hn}
dr{s>809:A,A}
nj{a<3614:dk,m<1828:bh,dn}
dqp{x>3865:A,a<2461:R,A}
brn{s<898:gd,m>2816:A,a>3580:R,R}
hhq{a<2868:stl,R}
hjc{a<231:R,s>2817:A,s<2334:R,R}
zb{a>3194:A,A}
hx{s>3290:cbx,a<3076:ccz,jh}
bf{x<1862:bfx,x>2056:lzl,m<3362:rdq,hsm}
trr{s<2831:R,m<1406:R,m>2129:A,A}
vb{x>3196:A,m>3164:A,R}
dv{m>2200:R,a<91:R,s<2940:A,R}
zkl{x<2284:A,s<1097:R,R}
gpj{x>1371:A,R}
cst{s<2470:R,s>2812:A,a<643:A,A}
bvc{s<1676:R,a>1007:A,s>1915:R,A}
rg{s<1279:A,a>3062:mmg,m>2345:lxk,A}
nzl{a>3748:xx,x>3703:R,kcr}
tj{a<3652:pc,zc}
kz{m>2638:A,R}
tm{x<3457:jvm,x<3680:dt,lqf}
jb{m>2515:njp,a>649:xzq,A}
scr{m>765:R,m>326:R,s>1957:zp,cnx}
jg{s>164:R,x<938:R,R}
vst{m>1299:vn,A}
vn{a>3343:R,s>2614:A,m>1576:A,A}
phn{s>1416:mmv,x>1736:R,A}
qrz{s>2030:qr,a<2576:xb,gmv}
cx{x<662:R,a>2695:A,a<2609:A,R}
pcn{x>2550:R,x>2521:R,s>3302:A,R}
kcr{a<3651:R,A}
cdh{x<3746:xsp,xj}
ml{a<1180:A,m<1386:R,s<2716:R,R}
qp{s>2644:krm,m<328:A,s<2515:R,rxh}
gd{m<2876:R,s>766:R,x<1002:R,R}
jkl{x>315:R,A}
kxf{a>1649:msb,s<2716:R,x>3763:A,hxc}
qdj{m<3458:A,R}
xc{a<1445:A,s<266:A,s<432:A,A}
pgb{x>3259:A,A}
js{m<1511:R,a<647:R,m<2107:R,A}
pc{a<3528:A,a<3599:R,A}
sf{s<1122:hr,R}
tnd{m>1346:lzm,a>253:cdh,bzb}
pz{a>275:dgc,x>3327:fq,x<3304:lvf,lxs}
zh{s>2600:R,s<2473:A,x>3528:A,R}
rgt{s<2437:R,a<1721:R,m>700:A,zhs}
qg{m>2511:shq,bn}
hh{x>1614:A,x>724:R,R}
xj{m>719:R,x>3904:A,R}
dlf{m>425:zgc,s<2770:A,zj}
zd{x>3507:ng,a<235:gkr,jzk}
lvf{x>3283:gq,x>3271:scf,s>3336:zmg,R}
bq{s<1389:ngv,x>3685:sz,a<2449:scr,nnp}
stl{s<3570:R,a>2493:R,m>1923:R,R}
jkd{x>2287:A,s>2347:R,R}
rrd{a<3633:R,s<866:A,R}
qdz{a<797:pt,x>3479:tzv,x<3219:nzq,pcl}
cnx{m<154:R,R}
vdr{m>821:ml,x<535:dlf,x<709:fdl,lxp}
zbx{x<2640:kbz,gzj}
ccz{m<2225:qnk,zk}
rhl{m>442:xzr,m>160:ptv,x<2137:R,xtb}
xx{s>1208:R,s>494:A,m>2651:R,A}
vqd{m<1086:A,R}
ft{x>3082:qf,A}
vfx{x>3029:A,a<354:R,s<1613:A,R}
vlz{x>1977:A,m<3547:R,A}
xb{x<1540:db,a<2432:kn,A}
vg{s>1485:pch,nj}
cps{a>347:A,s>1568:R,R}
ph{x>3890:R,x<3840:A,A}
nkb{a<1522:vb,s>3017:qm,a>1866:xr,sbk}
czv{a>2319:R,a<2240:A,m<2945:R,R}
rz{s<2233:sl,m<1762:htk,cqg}
jqh{m<2732:R,s<698:R,R}
cl{s<1943:pg,m>3067:jc,a>2554:zbf,btr}
qj{s>3571:R,s>3203:R,s>3040:R,A}
bnd{x>2730:ds,vzf}
dk{s>1261:qpx,hh}
jdh{s>885:zb,m<1833:A,a>3267:R,zqk}
zmd{x<1887:R,s>2777:R,A}
zr{x<1015:R,A}
mzz{a<2762:R,m<1597:R,A}
mt{x>2830:R,A}
fm{s>2355:A,a<640:A,A}
cgj{a>1400:A,A}
kkv{m>2421:R,m<1271:A,x>1461:A,R}
jnm{s>3324:tb,xf}
mmv{s<2291:A,m>589:R,R}
nzq{x<3080:R,x<3161:A,x>3193:pk,R}
qqn{a<3324:A,m>3316:A,s<353:R,A}
zg{m>3397:A,s<2871:hzq,R}
bx{s>2715:rbt,R}
rb{s<200:A,x<338:mq,a<445:A,hdg}
pzb{m<1052:mc,jnm}
pjj{s<3810:fdr,s<3840:A,x<2865:R,lqm}
vbb{x<3142:dsx,a>2939:kk,bd}
grs{a<1727:R,R}
xkk{s<839:sjl,x<3094:vfx,A}
gbn{s>374:A,R}
jd{m<1657:R,A}
vh{a>1414:grs,tv}
cfk{s>2792:R,R}
xkg{m>3401:qkp,rl}
qq{x>2570:A,s>1833:R,A}
rdr{x<3603:A,A}
vkp{x<2182:A,s>389:A,a<2301:R,R}
fjt{a>1688:jj,x>2307:A,s<2361:zkl,R}
gm{m>354:R,a<186:A,A}
prx{s>1512:A,s>1435:A,A}
gql{a<2017:A,x<1851:A,x>2090:A,R}
qnv{s<2438:A,mlj}
jc{s<2139:cfc,x>1303:R,jhn}
gvv{a>557:ck,zhj}
shq{a>2621:R,x<3568:A,m<3296:A,R}
hsr{x<777:fng,pvk}
cdv{m>1882:A,a<2476:A,A}
xgz{s>1149:jl,mhx}
zls{m>2280:tqt,pr}
nb{s<3894:R,m<3347:R,a<1324:R,A}
cf{x<3811:A,a<1522:R,a<1937:R,A}
hzq{a<3667:A,A}
cbx{a>3297:tn,cz}
krm{s<2733:A,R}
cc{s<663:fp,R}
phc{x<1414:A,m<2002:R,R}
kj{a>2731:R,R}
vnf{m<1410:ft,tjk}
mzl{s>3266:xd,kxf}
gh{s>1682:R,m<694:A,A}
nfb{m<2177:R,s<2536:A,a<136:A,A}
hl{x<1454:jpc,vkz}
dtt{a<3598:R,m<3110:A,R}
hf{s<3174:A,A}
mf{s>2274:R,A}
tfs{x>1223:dbt,a>841:A,R}
sz{s<1802:zf,s>2097:R,R}
hnx{x<1072:rf,x<1295:scx,a<1163:fh,lhd}
cqg{s<3299:vh,m<2775:hkn,a>1114:qb,rsq}
qln{a>1673:A,R}
pxc{a>2657:R,a>2381:A,A}
krh{s<291:R,m<2436:A,A}
dt{x>3600:srr,s>2701:R,s>2405:zh,mf}
tp{x<2421:A,a<498:R,a<740:R,R}
fl{m>1485:R,a>3798:A,A}
fdd{s<2907:R,a>639:R,tgn}
pk{s>3341:A,A}
xr{m<2706:xcr,m<3181:lfd,a>2002:qhj,A}
vtx{s>2820:R,m>1639:A,R}
hvk{m>1133:ph,R}
nr{m<2779:A,m<3519:R,R}
ff{s>646:A,a<407:A,R}
xzq{x<3589:A,R}
fh{x>1421:kkv,a<445:A,s>409:A,gpj}
vgk{s>907:R,m>1141:R,s<452:A,qx}
sx{a<650:A,s>2903:A,A}
xhz{x>1449:tmc,m>935:mzz,a<2771:A,vc}
th{m>2597:R,R}
nx{x<2811:np,a<2883:jx,x>2968:qt,gl}
zgc{x<233:A,s>2710:A,s>2651:A,R}
cfc{x>1027:A,A}
pvk{m>968:A,x<822:A,m<356:vq,ld}
ls{x<1420:A,A}
fn{a<1688:nb,s>3909:A,A}
qx{m<1058:R,m>1087:A,s<691:R,R}
tfn{s<568:krh,a<2462:jp,x>3504:dnh,R}
nv{m<2414:ll,m<3178:R,s<1651:A,R}
txm{a>3840:rdr,x<3493:sf,m<1870:xxx,nzl}
fs{m>370:A,x>656:A,A}
lzb{x<3553:A,a<260:A,R}
scx{m<2306:R,mxt}
pm{x>1758:A,R}
rs{m<837:rt,a<3366:A,kf}
vt{m<1064:A,s<373:R,m<1360:R,R}
jst{a>676:qdz,s>3154:cj,a>616:vls,tm}
gz{s<2338:A,s<2438:R,A}
jvm{a<572:A,x>3220:R,bcx}
mhq{m<1439:A,s>2440:R,R}
fns{s<2559:zn,a<2754:kms,R}
jk{x<171:htv,m>942:bvr,hcb}
zbf{a>2970:A,R}
cz{x<1546:lf,s<3744:hhq,s>3856:tk,pjj}
rf{s>369:xrr,jg}
vbm{m<2344:sd,x<1876:zg,a>3666:lj,zpx}
xv{a<2788:cdl,x<566:gbn,vpf}
zp{m>201:A,m<128:R,x<3340:R,A}
lqm{m>1799:A,x>3463:A,s<3848:R,R}
jpc{x>531:A,x>317:R,m<2632:A,R}
lzm{a<231:R,s>3024:R,pmv}
qlt{x<3440:cst,bc}
lt{a>2699:rgx,s>370:bdc,s>149:qzq,xt}
mxr{s<485:R,a>1675:R,a<1193:R,A}
fds{m<2448:nz,m<3000:xkk,dr}
ll{x<1132:A,x<1280:R,a<1573:A,R}
kms{s<2765:A,x<2309:A,m<3173:A,A}
gb{s<3891:R,R}
dgc{a<411:A,a<452:jd,a>498:A,R}
tb{s>3700:R,A}
lr{a>879:jhf,s<2053:cht,a<524:nrf,jst}
rmb{m>1101:mhq,x>1215:A,R}
jkc{m<1158:qgh,A}
tz{s>3126:R,m>419:A,A}
sbk{x>3333:R,R}
zhj{x>1246:R,s<1716:R,m<1604:R,R}
vjj{x<3215:R,m>1566:dxd,x<3278:hj,A}
hr{s>540:R,m<2033:R,s<282:R,A}
cj{s<3589:fx,a>597:jb,cdt}
gn{a<1595:R,x<3715:A,R}
hcs{x>172:vt,s>635:A,cp}
nqg{s>3402:A,R}
vzf{a<1644:qq,R}
tv{m<3244:A,R}
dmq{a>3689:A,a<3475:A,s<1971:lc,A}
nrf{x>3585:tnd,x<3263:vnf,x<3447:pz,zd}
xld{s<1715:A,A}
rdq{a<516:R,A}
bqr{x<3847:R,R}
rfz{s<1041:mxr,a>1573:R,a>1257:dlj,bvc}
nzk{x<3557:A,R}
jf{s<3564:A,x<974:R,s>3672:R,A}
nz{s<1151:rqr,cps}
pcl{s<2881:zz,A}
rh{a>449:R,m<2584:R,xps}
pb{a<2666:A,a<2952:R,A}
nrd{s>1704:R,s>1330:prx,a<1711:A,zsg}
qtd{s<612:ts,m<2565:lfs,m>3280:kh,brn}
pkf{x<3271:R,a>3042:A,s>205:A,A}
fkb{s>865:R,s>823:A,A}
dtk{a>1109:rxs,vrt}
gmv{s<1859:R,R}
gl{m<1353:vmq,x>2904:nr,x<2843:jxq,xvg}
nnp{s>1853:R,s>1577:gh,x<3502:qn,A}
vv{x<2224:A,m<1830:R,A}
vx{x<2424:vqd,s>2665:kq,tr}
zn{s<2449:R,x>2073:A,R}
tzv{a>840:vtx,m>2659:A,gqn}
czr{m>2158:gb,R}
pch{x>902:gtk,x<487:dmq,s>1811:cq,lp}
fx{s>3350:ksm,x>3372:A,njl}
dd{a>3392:A,a>3268:A,x<645:A,R}
xps{x>1814:R,R}
mcj{x>2067:R,A}
jds{a>3112:R,A}
rxh{a>3347:A,s>2592:R,m<507:A,R}
bd{m<1603:bq,a>2541:sb,rcm}
lpl{x<1753:R,a<1523:A,R}
ng{s>2948:nzk,lzb}
pt{m<2493:trr,m<3406:cfk,a>741:A,R}
mn{x<1933:phn,x>2246:fjt,rhl}
tl{x>2685:A,x>2615:A,R}
bpd{s>2076:bk,m<1552:vgk,x<1995:ql,vv}
fdl{x>616:fs,a>810:R,x<564:pf,xl}
cg{a<2496:vkp,A}
jhn{m>3393:R,m<3279:R,A}
sb{a<2675:qg,s>1560:trh,a<2768:djq,bb}
tqt{s<1353:A,R}
scf{a<158:R,R}
vkz{a>2439:R,A}
hm{a>431:ghl,R}
rqg{m>2575:R,x>556:A,a>2512:A,A}
htv{m<997:R,m<1506:R,s>542:R,R}
hcb{m>466:A,a<2797:A,m>284:R,R}
rc{x>3561:R,a<3753:A,R}
nrz{a<653:A,x>2483:R,mmb}
rsq{s>3533:gcp,qsx}
cdl{x<594:clm,a<2563:gqt,x<633:A,cx}
ngv{s>869:pxc,s>544:bcb,x<3499:A,A}
vbq{s>2895:ghp,a<3253:vx,m<696:qp,vst}
db{a<2416:A,x>775:R,m>1331:A,A}
rgd{m<3005:zfk,s>584:gn,m<3486:xc,kfj}
nh{a>3582:R,m>1171:A,s>2047:A,A}
kzx{m>2108:rr,x>3758:hvk,rs}
xm{a<2478:R,s<286:R,m>2319:R,R}
gc{m>1801:R,s<426:A,A}
jmk{m<1508:dm,m<2671:A,a<1026:A,ctk}
txc{x<1334:tt,A}
tjk{x<3091:R,a<312:R,A}
rgx{a<2981:R,s<685:A,jds}
srr{s>2719:R,m>2597:R,x<3647:R,A}
dn{a<3852:R,xz}
zc{s>3736:R,x>2493:A,x>1496:R,A}
qd{s>2041:R,x>1781:A,R}
jh{a>3404:vbm,m<1726:vbq,lck}
gqn{m<1677:R,x<3729:A,m<2060:A,A}
zq{a>1828:A,A}
bb{a<2825:R,s<1018:R,R}
qmb{x<2466:tp,s>573:xn,R}
fq{a<148:xg,m>1807:hjc,A}
kp{m>841:bx,tfs}
xl{x>592:A,s<2710:A,a<480:R,R}
cht{x<3317:fds,zqq}
rm{s>708:R,R}
zx{m>2501:jf,a<954:R,m<2318:A,A}
rqr{m<1218:A,m<1825:R,R}
lkq{a<983:gvv,nv}
xrr{x>961:A,x<901:R,R}
qh{x>2768:A,a<3300:R,m<3254:A,A}
vmq{a<3627:R,a<3872:R,a<3943:A,R}
qm{m<3328:sc,s<3556:A,m>3589:sp,qdj}
xd{m<2717:trb,m<3353:sj,nxs}
qrc{x<1219:A,R}
fp{m>2282:A,R}
svc{x<1578:A,x<2619:A,s>3904:kz,dtt}
bfx{a<463:A,a<685:R,m>3342:R,R}
dgj{m>3248:R,a>373:A,x<2392:A,R}
hb{m<468:R,s<3450:R,A}
gzv{m<2066:A,x>3762:A,a<559:A,A}
tfk{a>1442:R,x>3478:R,A}
djq{m>2471:kj,R}
zfk{m<2582:R,s<665:R,R}
dnh{m>2463:R,m>2120:A,A}
qgh{x<731:R,m<573:R,s<2388:A,R}
dx{m<2824:cc,df}
ttt{a>3204:qtd,x<893:lt,dx}
kfj{m>3806:R,m<3668:R,s>340:A,R}
rrq{s>458:A,x>3569:R,R}
clm{s>410:A,s>167:R,A}
xt{m<3146:A,R}
vsr{s<3421:A,A}
np{x<2718:R,m<2531:A,qh}
lxk{x<888:R,a>2893:R,A}
br{a<210:A,A}
cb{m<843:R,a>120:A,x>3826:A,R}
bdc{a>2355:rqg,a<2266:jqh,a>2320:td,mgl}
njp{a<633:A,R}
bk{a>1796:gql,m<1358:zmd,m>1825:A,A}
tmc{x>3023:R,R}
dbq{s<741:R,A}
tk{s<3950:A,s<3970:mt,x<2569:R,R}
cq{m>2542:djk,nh}`

const input = `{x=797,m=1039,a=391,s=528}
{x=2289,m=300,a=1584,s=196}
{x=954,m=93,a=607,s=1988}
{x=1432,m=395,a=965,s=639}
{x=1251,m=3236,a=34,s=579}
{x=124,m=32,a=2829,s=2331}
{x=1322,m=36,a=848,s=1279}
{x=144,m=3450,a=3,s=3497}
{x=1490,m=2205,a=1062,s=43}
{x=82,m=581,a=917,s=64}
{x=426,m=360,a=406,s=2325}
{x=903,m=1148,a=2024,s=1317}
{x=1244,m=1284,a=503,s=959}
{x=822,m=2553,a=132,s=229}
{x=1712,m=2155,a=2103,s=1734}
{x=498,m=1757,a=2591,s=1737}
{x=2085,m=440,a=105,s=1625}
{x=543,m=1928,a=217,s=24}
{x=894,m=110,a=99,s=473}
{x=377,m=1927,a=2263,s=3432}
{x=178,m=9,a=3305,s=2256}
{x=34,m=277,a=471,s=639}
{x=403,m=154,a=1871,s=1139}
{x=662,m=1151,a=37,s=526}
{x=1550,m=2909,a=481,s=1614}
{x=272,m=1674,a=67,s=2356}
{x=152,m=211,a=1420,s=1256}
{x=1113,m=1378,a=113,s=1930}
{x=2040,m=2575,a=1092,s=531}
{x=693,m=772,a=73,s=843}
{x=1,m=1347,a=834,s=4}
{x=1270,m=1653,a=1823,s=1582}
{x=304,m=385,a=243,s=2712}
{x=1619,m=650,a=771,s=596}
{x=2324,m=873,a=67,s=43}
{x=1210,m=487,a=917,s=1167}
{x=1272,m=173,a=1613,s=731}
{x=188,m=1010,a=184,s=1612}
{x=1315,m=2730,a=28,s=2682}
{x=1899,m=966,a=312,s=138}
{x=458,m=1260,a=1333,s=1023}
{x=399,m=2098,a=122,s=103}
{x=290,m=232,a=333,s=2521}
{x=853,m=624,a=189,s=760}
{x=271,m=652,a=1882,s=90}
{x=75,m=389,a=524,s=758}
{x=800,m=416,a=706,s=717}
{x=1188,m=307,a=51,s=33}
{x=1201,m=888,a=3423,s=650}
{x=2964,m=978,a=2323,s=75}
{x=949,m=536,a=758,s=2114}
{x=1226,m=3552,a=790,s=13}
{x=100,m=48,a=255,s=194}
{x=1533,m=2176,a=1889,s=1173}
{x=109,m=1433,a=345,s=176}
{x=1772,m=81,a=1833,s=745}
{x=482,m=277,a=815,s=341}
{x=94,m=39,a=1093,s=67}
{x=1761,m=2080,a=18,s=8}
{x=514,m=554,a=1270,s=354}
{x=604,m=378,a=47,s=879}
{x=799,m=39,a=186,s=161}
{x=448,m=2282,a=98,s=88}
{x=133,m=363,a=179,s=622}
{x=253,m=34,a=823,s=52}
{x=144,m=1413,a=267,s=3729}
{x=619,m=1895,a=611,s=170}
{x=2301,m=3461,a=2609,s=529}
{x=397,m=69,a=2421,s=117}
{x=15,m=1333,a=94,s=46}
{x=569,m=1540,a=1609,s=685}
{x=1975,m=887,a=3087,s=203}
{x=1690,m=2970,a=1630,s=899}
{x=2028,m=17,a=391,s=2716}
{x=237,m=984,a=3185,s=1656}
{x=1225,m=2146,a=845,s=1456}
{x=1916,m=694,a=687,s=21}
{x=1047,m=752,a=2041,s=1562}
{x=2358,m=673,a=556,s=1234}
{x=637,m=241,a=877,s=1078}
{x=2434,m=109,a=132,s=2856}
{x=1530,m=155,a=572,s=362}
{x=779,m=172,a=1967,s=1212}
{x=1209,m=1637,a=1105,s=93}
{x=1511,m=744,a=2089,s=638}
{x=1564,m=240,a=284,s=927}
{x=78,m=261,a=565,s=934}
{x=30,m=41,a=3435,s=15}
{x=117,m=900,a=659,s=125}
{x=844,m=2713,a=331,s=14}
{x=880,m=31,a=222,s=442}
{x=604,m=1172,a=908,s=855}
{x=1332,m=2063,a=1210,s=454}
{x=1277,m=32,a=2631,s=95}
{x=444,m=460,a=23,s=1595}
{x=1317,m=1377,a=448,s=284}
{x=520,m=1540,a=196,s=1057}
{x=2225,m=268,a=981,s=1980}
{x=892,m=1,a=457,s=3123}
{x=335,m=2375,a=1949,s=305}
{x=789,m=374,a=201,s=596}
{x=3058,m=1588,a=952,s=2199}
{x=1773,m=56,a=110,s=2339}
{x=1009,m=602,a=616,s=1043}
{x=66,m=442,a=161,s=108}
{x=249,m=1400,a=3486,s=559}
{x=44,m=341,a=214,s=1060}
{x=3973,m=353,a=953,s=817}
{x=1456,m=2149,a=2871,s=2116}
{x=2146,m=1018,a=940,s=1005}
{x=373,m=884,a=1013,s=32}
{x=99,m=949,a=2104,s=185}
{x=615,m=96,a=821,s=1999}
{x=331,m=652,a=56,s=1513}
{x=885,m=1527,a=1236,s=1244}
{x=813,m=1792,a=2002,s=587}
{x=5,m=53,a=460,s=2283}
{x=2284,m=1240,a=2377,s=138}
{x=97,m=3384,a=2243,s=1333}
{x=93,m=1292,a=250,s=766}
{x=585,m=2922,a=1347,s=1963}
{x=262,m=467,a=334,s=33}
{x=586,m=247,a=1165,s=1294}
{x=5,m=498,a=86,s=435}
{x=2642,m=905,a=482,s=45}
{x=453,m=859,a=402,s=2222}
{x=697,m=1885,a=496,s=1298}
{x=2760,m=911,a=1904,s=274}
{x=7,m=685,a=139,s=127}
{x=2233,m=115,a=981,s=286}
{x=1077,m=1875,a=953,s=527}
{x=704,m=2593,a=849,s=1436}
{x=1707,m=871,a=840,s=2295}
{x=76,m=938,a=329,s=640}
{x=1976,m=878,a=1419,s=875}
{x=7,m=1132,a=951,s=1005}
{x=200,m=2098,a=1626,s=1339}
{x=3487,m=1635,a=902,s=841}
{x=1293,m=37,a=82,s=448}
{x=1086,m=150,a=1785,s=422}
{x=671,m=74,a=927,s=1992}
{x=627,m=1033,a=1015,s=495}
{x=1954,m=734,a=1775,s=410}
{x=747,m=3196,a=2323,s=465}
{x=6,m=1904,a=784,s=69}
{x=2616,m=3470,a=1547,s=256}
{x=1225,m=1221,a=85,s=612}
{x=668,m=2019,a=47,s=854}
{x=807,m=2700,a=1365,s=2155}
{x=366,m=816,a=620,s=675}
{x=460,m=1805,a=2487,s=1327}
{x=125,m=1991,a=1890,s=13}
{x=3074,m=1445,a=13,s=1301}
{x=508,m=735,a=241,s=1462}
{x=861,m=607,a=968,s=1766}
{x=1280,m=618,a=221,s=1449}
{x=2526,m=423,a=791,s=1207}
{x=1951,m=1166,a=1301,s=856}
{x=957,m=15,a=1388,s=1640}
{x=1113,m=784,a=1565,s=2578}
{x=118,m=2209,a=1819,s=2254}
{x=1252,m=438,a=1822,s=1327}
{x=1979,m=384,a=1869,s=1132}
{x=19,m=461,a=888,s=447}
{x=2,m=189,a=59,s=2047}
{x=1302,m=587,a=2644,s=580}
{x=1054,m=145,a=1146,s=222}
{x=1199,m=2666,a=156,s=31}
{x=1838,m=888,a=1461,s=549}
{x=414,m=296,a=1849,s=362}
{x=2567,m=2441,a=59,s=1317}
{x=696,m=2054,a=1621,s=499}
{x=1592,m=1069,a=880,s=3390}
{x=177,m=1213,a=25,s=2376}
{x=69,m=119,a=789,s=241}
{x=1054,m=710,a=82,s=48}
{x=112,m=1415,a=1363,s=920}
{x=567,m=14,a=1106,s=142}
{x=1596,m=230,a=1413,s=100}
{x=731,m=2270,a=630,s=2538}
{x=620,m=434,a=595,s=1779}
{x=759,m=77,a=1502,s=211}
{x=310,m=1377,a=1569,s=523}
{x=2693,m=2324,a=2075,s=1136}
{x=13,m=759,a=1326,s=468}
{x=399,m=517,a=1395,s=538}
{x=738,m=732,a=579,s=97}
{x=1563,m=799,a=1145,s=326}
{x=993,m=179,a=302,s=233}
{x=1740,m=2864,a=1048,s=592}
{x=854,m=1492,a=14,s=2628}
{x=1535,m=2408,a=773,s=1822}
{x=1170,m=328,a=645,s=303}
{x=365,m=460,a=1308,s=1575}
{x=1307,m=1140,a=181,s=989}
{x=428,m=66,a=1759,s=261}
{x=2199,m=528,a=2296,s=1949}
{x=10,m=887,a=1,s=44}
{x=300,m=1240,a=928,s=610}
{x=1487,m=1772,a=140,s=817}`

const result = 389114

function wordGrouping(words) {
    let wordsArray = words.split(",");
    let resultArr = [];

    for (let i = 0; i < wordsArray.length; i++) {
        let tmp = [];
        for(let j = 0; j < wordsArray.length; j++) {
            if ((JSON.stringify(wordsArray[i].split("").sort()) === JSON.stringify(wordsArray[j].split("").sort()))) {
                tmp.push(wordsArray[j]);
            }
        }
        tmp = tmp.filter(function (item, pos) {
            return tmp.indexOf(item) == pos;
        })
        resultArr.push(tmp.join("-"));
    }
    resultArr = resultArr.filter(function (item, pos) {
        return resultArr.indexOf(item) == pos;
    })
    console.log(resultArr);
}
wordGrouping("VMRCO,VORCM,MCRTOX,ZXTAC,XZATC,XMTCOR,OXVS,AZTXC,VXOS,RZAT,MRCOTX,SVXO,TRAZ,ZTAR,MVOCR");

function findFee(t) {
    let hour = 0;
    let min = 0;
    let total = 0;
    if (t.split(" hours")[1] === undefined) {
        min = parseInt(t.split(" minutes")[0]);
    } else{
        hour = parseInt(t.split(" hours")[0]);
        min = parseInt(t.split(" hours")[1].split(" minutes")[0]);
    }
    total = 60*hour + min;
    if(total < 30){
         console.log("Free");
    } else if(total <= 180){
        console.log(Math.ceil(total/60)*10 + " bath");
    } else if(total <= 360){
        console.log(Math.ceil(total/60)*20 + " bath");
    } else if(total > 360){
        console.log(Math.ceil(total/1440)*200 + " bath");
    }
}
findFee("2 hours 10 minutes");

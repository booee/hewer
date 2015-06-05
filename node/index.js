var minimist = require('minimist');
var fs = require('fs');

// parameters
var args = minimist(process.argv.slice(2));
var fileName = args._[0];
var key = args.k || args.key;




// parse file
if(!fileName) {
    console.log('Provide a file name, fool!');
    return;
}

var fileContent = fs.readFileSync(fileName).toString().split("\n");
var rows = [];

for(var i = 0, len = fileContent.length; i < len; i++) {
    var rawString = fileContent[i];
    if(rawString) {
        var obj = JSON.parse(rawString);
        rows.push(obj);
    }
}




// refine
var refinedRows = [];
var extraProperties = [];
var allNumbers = true;
for(var i = 0, len = rows.length; i < len; i++) {
    var row = rows[i];
    var refined = nestedGet(row, key);
    if(!isNaN(refined)) {
        refined = Number(refined);
    } else {
        allNumbers = false;

        if(refined instanceof Object) {
            tempKeys = Object.keys(refined);
            for(var k = 0, kLen = tempKeys.length; k < kLen; k++) {
                var tempKey = tempKeys[k];
                if(extraProperties.indexOf(tempKey) < 0) extraProperties.push(tempKey);
            }
        }
    }

    refinedRows.push(refined);
}

function nestedGet(row, valueName) {
    if(!key) return row;

    var keys = valueName.split('.');
    var holder = row;

    for(var i = 0, len = keys.length; i < len; i++) {
        if(holder === undefined) {
            return undefined;
        }

        holder = holder[keys[i]];
    }

    return holder;
}




// build analysis & display
var analysis = {};
if(key) analysis.key = key;

analysis.entries = refinedRows.length;

if(allNumbers) {
    analysis.max = Math.max.apply(this, refinedRows);
    analysis.min = Math.min.apply(this, refinedRows);

    var sum = refinedRows.reduce(function(a, b) { return a + b; });
    analysis.avg = sum / analysis.entries;
} else {
    if(extraProperties.length) analysis.properties = extraProperties;
}

console.log(analysis);

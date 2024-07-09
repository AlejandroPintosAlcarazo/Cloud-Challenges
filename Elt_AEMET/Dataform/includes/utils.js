function DMS_TO_DD(dms) {
    var direction = dms.slice(-1);
    var degrees = parseFloat(dms.slice(0, dms.length - 5));
    var minutes = parseFloat(dms.slice(dms.length - 5, dms.length - 3));
    var seconds = parseFloat(dms.slice(dms.length - 3, dms.length - 1));
    var dd = degrees + (minutes / 60) + (seconds / 3600);
    if (direction == 'S' || direction == 'W') {
      dd = dd * -1;
    }
    return dd;
  }
  
  module.exports = { DMS_TO_DD };
  
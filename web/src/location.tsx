import React, { Component } from "react";

function errors(err) {
  // error ???
}
export var getPosition = function(cb) {
  if (navigator.geolocation) {
    navigator.permissions
      .query({ name: "geolocation" })
      .then(function (result) {
        if (result.state === "granted") {
          //If granted then you can directly call your function here
          navigator.geolocation.getCurrentPosition(cb);
        } else if (result.state === "prompt") {
          navigator.geolocation.getCurrentPosition(cb, errors, {enableHighAccuracy: true,timeout: 5000,maximumAge: 0});
        } else if (result.state === "denied") {
          //If denied then you have to show instructions to enable location
        }
        result.onchange = function () {
          console.log(result.state);
        };
      });
  } else {
    // error ???
  }
}; 

export default class GeoLocation extends Component {}
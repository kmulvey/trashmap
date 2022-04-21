import * as React from 'react';
import {useState, useCallback} from 'react';
import {render} from 'react-dom';
import Map from 'react-map-gl';

import DrawControl from './draw-control';
import ControlPanel from './control-panel';
import GeoLocation, { getPosition } from './location';

const TOKEN = ''; // Set your mapbox token here

export default function App() {
  const [features, setFeatures] = useState({});

  const onUpdate = useCallback(e => {
    setFeatures(currFeatures => {
      const newFeatures = {...currFeatures};
      for (const f of e.features) {
        newFeatures[f.id] = f;
      }
      return newFeatures;
    });
  }, []);

  const onDelete = useCallback(e => {
    setFeatures(currFeatures => {
      const newFeatures = {...currFeatures};
      for (const f of e.features) {
        delete newFeatures[f.id];
      }
      return newFeatures;
    });
  }, []);

  var lat;
  var long; 
  getPosition(function(pos){
    lat = pos.coords.latitude;
    long = pos.coords.longitude;
  });

  return (
    <>
      <Map
        initialViewState={{
          longitude: long,
          latitude: lat,
          zoom: 12
        }}
        mapStyle="mapbox://styles/mapbox/streets-v9"
        mapboxAccessToken={TOKEN}
      >
        <DrawControl
          position="top-left"
          displayControlsDefault={false}
          controls={{
            polygon: true,
            trash: true
          }}
          defaultMode="draw_polygon"
          onCreate={onUpdate}
          onUpdate={onUpdate}
          onDelete={onDelete}
        />
      </Map>
      <ControlPanel polygons={Object.values(features)} />
      <GeoLocation />
    </>
  );
}

export function renderToDom(container) {
  render(<App />, container);
}

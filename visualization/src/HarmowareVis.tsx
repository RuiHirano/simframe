import React, { useContext } from 'react';
import { useLineMapLayer, useHarmowareVis, useMovesLayer, useDepotsLayer, useController } from 'harmoware-vis-utility-hooks'
const MAPBOX_TOKEN = 'pk.eyJ1IjoicnVpaGlyYW5vIiwiYSI6ImNqdmc0bXJ0dTAzZDYzem5vMmk0ejQ0engifQ.3k045idIb4JNvawjppzqZA'


const createMovesData = () => {
  const step = 100
  const moves_num = 100
  const time = Math.floor(new Date().getTime() / 1000)
  const data = []

  for (let i = 0; i < step; i++) {
    for (let k = 0; k < moves_num; k++) {
      data.push({
        type: k,
        operation: [
          {
            elapsedtime: time + i,
            position: [136.979052, 35.152912, 0],
            color: [0, 255, 0],
          }
        ]
      })
    }
  }
  console.log('data', data)
  return data
}

const data = createMovesData()

export const HarmowareVisComponent: React.FC = () => {

  const { setMovesLayer } = useMovesLayer({
    data: data,
  })

  const { setLineMapLayer } = useLineMapLayer({
    data: [  // line source & target
      {
        "sourcePosition": [136.978052, 35.152912, 0], // line start position (long,Lati,height)
        "targetPosition": [136.981445, 35.157597, 0], // line end position (long,Lati,height)
      },
      {
        "sourcePosition": [136.979052, 35.152912, 0], // line start position (long,Lati,height)
        "targetPosition": [136.981445, 35.157597, 0], // line end position (long,Lati,height)
      },
      {
        "path": [[136.977052, 35.152912, 0], // line path position (long,Lati,height)
        [136.978052, 35.152912, 0]],
        "dash ": [5, 2], // line pattern
      },
      {
        "polygon": [[136.977052, 35.152812, 0], // polygon path position (long,Lati,height)
        [136.976052, 35.152912, 0], [136.975052, 35.152312, 0]],
        "elevation": 10, // 3-D object height
      },
      {
        "coordinates": [[136.977752, 35.152212, 0], // coordinates path position (long,Lati,height)
        [136.977852, 35.152512, 0]],
      },
    ]
  })

  const { renderHarmowareVis } = useHarmowareVis({
    mapbox_token: MAPBOX_TOKEN,
    layers: [
      setMovesLayer(),
      setLineMapLayer()
    ],
    viewport: {
      longitude: 136.9831702,
      latitude: 35.1562909,
      width: window.screen.width,
      height: window.screen.height,
      zoom: 16
    }
  })

  const { renderController } = useController()


  return (
    <div className="App">
      {renderController()}
      {renderHarmowareVis()}
    </div>
  );
}
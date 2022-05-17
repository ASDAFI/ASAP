import { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import { Loading } from "./Loading";
import { ResponsiveBump } from '@nivo/bump'

import "./WaterLog.css";
const styles = {
    fontFamily: "sans-serif",
    textAlign: "center"
  };
  function getDataFrameBySerial(params) {
      
  }
  const cars = new Array();
  for (let index = 0; index < 30; index++) {
    cars.push({"x":index , "y":Math.floor(Math.random() * 10) });
    
  }
  const data = [
    {
      "id": "Serie 1",
      "data": [
        {
          "x": 2000,
          "y": 2
        },
        {
          "x": 2001,
          "y": 9
        },
        {
          "x": 2002,
          "y": 2
        },
        {
          "x": 2003,
          "y": 4
        },
        {
          "x": 2004,
          "y": 6
        }
      ]
    }
  ];
export const ReportLog = () => {
  const cookies = new Cookies();
  const [waterLogs, setWaterLogs] = useState({});
  useEffect(() => {
    fetch("http://asap.carriot.ir:8000/device/logs/serial", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + cookies.get("token"),
        "content-type": "application/json",
      },
    }).then((res) => {
      res.json().then((logs) => {
        setWaterLogs(logs);
      });
    });
  }, []);

  return (
    <div
      style={{
        backgroundColor: "#fff",
        width: "70%",
        borderRadius: 14,
        marginTop: 14,
        paddingLeft: 14,
      }}
    >
      <h3 style={{ color: "#00aaff" }}>Humidity Report</h3>
      <div style={styles}>
    <h1>Nivo basic demo</h1>
    <div style={{ height: "400px" }}>
    <ResponsiveBump
        data={data}
        xPadding={0.3}
        xOuterPadding={0.2}
        yOuterPadding={0.6}
        colors={{ scheme: 'brown_blueGreen' }}
        lineWidth={3}
        activeLineWidth={6}
        inactiveLineWidth={3}
        inactiveOpacity={1}
        startLabelTextColor={{ from: 'color', modifiers: [] }}
        pointSize={0}
        activePointSize={20}
        inactivePointSize={8}
        pointColor={{ theme: 'background' }}
        pointBorderWidth={3}
        activePointBorderWidth={3}
        pointBorderColor={{ from: 'serie.color' }}
        axisTop={{
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
            legend: '',
            legendPosition: 'middle',
            legendOffset: -36
        }}
        axisBottom={{
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
            legend: '',
            legendPosition: 'middle',
            legendOffset: 32
        }}
        axisLeft={{
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
            legend: 'ranking',
            legendPosition: 'middle',
            legendOffset: -40
        }}
        margin={{ top: 40, right: 100, bottom: 40, left: 60 }}
        axisRight={null}
    />   </div>
  </div>
      
    </div>
  );
};

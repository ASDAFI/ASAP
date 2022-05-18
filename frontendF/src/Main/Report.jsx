import { useEffect, useState } from "react";
import Cookies from "universal-cookie";
import { FInput2, TimeInput } from "../Login/FInput";
import { Loading } from "./Loading";
import { ResponsiveBump } from '@nivo/bump'

import "./WaterLog";

export const Report = ({ devices}) => {
  const cookies = new Cookies();
  const [reports, setreports] = useState({});
  const [startTime, setStartTime] = useState("");
  const [endTime, setEndTime] = useState("");
  const [ufCheck, setUFCheck] = useState({});
  const [logStatus, setLogStatus] = useState(0);
  const [chartData, setChartData] = useState({})


  useEffect(() => {
    fetch("http://asap.carriot.ir:8000/device/logs/serial", {
      method: "POST",
      headers: {
        Authorization: "Bearer " + cookies.get("token"),
        "content-type": "application/json",
      },
    }).then((res) => {
      setLogStatus(res.status);
      res.json().then((logs) => {
        setreports(logs);
        console.log("logs", logs);
      });
    });
  }, [ufCheck]);

  return (
    <div
      style={{
        backgroundColor: "#fff",
        width: "70%",
        borderRadius: 14,
        marginTop: 14,
        paddingLeft: 14,
        paddingRight: 14,
      }}
    >
      <h3
        style={{
          color: "#000000",
          borderBottomStyle: "solid",
          borderBottomColor: "#ccc",
          borderBottomWidth: 1,
          paddingBottom: 14,
          fontSize: "30px"
        }}
      >
        Report
      </h3>
      
      <h4 style={{ marginTop: 14, marginBottom: 2, fontSize: "22px" }}>Time Frame</h4>
      <div
        style={{ display: "flex", flexDirection: "row", alignItems: "center" }}
      >
        <h5 style={{ marginRight: "10px",
            fontSize: "15px"}}> choose Device :</h5>
        <select id="reportDevice" className="reportSelect" style={{ padding: "8px", fontSize: "16px"}}>
          {devices.devices ? (
            devices.devices.map((item) => {
              //console.log(item.device_serial);
              return (
                <option value={item.device_serial}>{item.device_serial}</option>
              );
            })
          ) : (
            <option value="">no device</option>
          )}
        </select>
        
      </div>
      <div
        style={{
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          fontSize: "15px"
        }}
      >
    
        <h5 style={{fontSize: "16px"}}>Start Time</h5>
        <TimeInput setTime={setStartTime} />
        <h5 style={{fontSize: "16px"}}>End Time</h5>
        <TimeInput setTime={setEndTime} />
      </div>
      <button
          className="chooseDeviceButton"
          onClick={() => {
            if (document.getElementById("reportDevice").value.length > 0) {
              if (startTime.length > 0 && endTime.length > 0) {
                fetch(
                  "http://asap.carriot.ir:8000/device/logs/serial",

                  {
                    method: "POST",
                    headers: {
                      Authorization: "Bearer " + cookies.get("token"),
                      "content-type": "application/json",
                    },
                    body: JSON.stringify({
                      device_serial:
                        document.getElementById("reportDevice").value,
                        start_date: new Date(startTime).toISOString(),
                        end_date: new Date(endTime).toISOString(),
                      step: 1,
                    }),
                  }
                ).then(res => res.json()) .then(data =>  {
                  
                    
                    let response = data;
                    
                    

                    let device_serial = response.device_serial
        
                  
                   
                    let logs = response.device_logs;
                    
                    
                    // fix undifined
                    
                    let x = {}
                    let y = {}
                    for (let i = 0; i < logs.length; i++) {
                      const time = new Date(logs[i].device_time).toLocaleDateString() 
                      x[i] = time
                      y[i] = logs[i].humidity
                    }
                    console.log(x)
                    console.log(y)
                    setChartData({
                        id: device_serial,
                        data: [{
                            x: x,
                            y: y
                        }
                        ]
                    }).catch((err)=>{
                        console.log("err", err)
                    })
                    
                    return(<div
                        style={{
                          backgroundColor: "#fff",
                          width: "70%",
                          borderRadius: 14,
                          marginTop: 14,
                          paddingLeft: 14,
                        }}
                      >
                        <h3 style={{ color: "#00aaff" }}>Humidity Report</h3>
                        <div style={{fontFamily: "sans-serif",
    textAlign: "center"}}>
                      <h1>Nivo basic demo</h1>
                      <div style={{ height: "400px" }}>
                      <ResponsiveBump
                          data={chartData}
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
                        
                      </div>);
                    //alert("report added successfully");

                    
                  
                  //console.log(res);
                });
              } else {
                alert("please fill the data");
              }
            } else {
              alert(
                "you have no devices\nyou can create devices from the menu"
              );
            }
          }}
        >
          show
        </button>
    </div>
  );
};

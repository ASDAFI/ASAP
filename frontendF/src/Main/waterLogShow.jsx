import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import Cookies from "universal-cookie";

import { Menu } from "./Menu";
import { WaterLog } from "./WaterLog";
import { Loading } from "./Loading";
import { ReportLog } from "./chart";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";

const WaterS = () => {
  const navigate = useNavigate();
  const cookies = new Cookies();
  const [user, setUser] = useState({});
  const [devices, setDevices] = useState([]);


  useEffect(() => {
    if (!cookies.get("token")) {
      navigate("/login");
    } else {
      fetch("http://asap.carriot.ir:8000/user", {
        method: "GET",
        headers: {
          Authorization: "Bearer " + cookies.get("token"),
          "content-type": "application/json",
        },
      }).then((res) => {
        if (res.status !== 200) {
          alert("an error occurred\nERROR Code : " + res.status);
        } else {
          res.json().then((rUser) => {
            setUser(rUser);
          });
        }
      });

    }
  }, []);

  return (
    <div>
      <img
        src={require("../Assets/loginBG.jpg")}
        alt="bg"
        style={{
          width: "100%",
          position: "absolute",
          left: 0,
          top: 0,
          overflow: "hidden",
          zIndex: -1,
        }}
      />

      <div
        style={{
          backgroundColor: "#ffffff70",
          borderRadius: 80,
          width: "70vw",
          // height: 400,
          paddingBottom: 40,
          marginTop: "8%",
          overflow: "hidden",
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        {!user.username ? (
          <Loading />
        ) : (
          <>
            <div
              style={{ display: "flex", flexDirection: "column", width: "90%" }}
            >
              <div
                style={{
                  display: "flex",
                  flexDirection: "row",
                  alignItems: "center",
                }}
              >
                <img
                  src={require("../Assets/flowerIcon.png")}
                  alt="icon"
                  style={{ width: 40, height: 40 }}
                />
                <h1>Farm Manager</h1>
              </div>
              <div
                style={{
                  borderBottomWidth: 2,
                  borderColor: "#000",
                  borderBottomStyle: "solid",
                  display: "flex",
                  flexDirection: "row",
                  alignItems: "center",
                  justifyContent: "space-between",
                }}
              >
                <p style={{ fontSize: 18 }}>
                  {user.first_name +
                    " " +
                    user.last_name +
                    " ( " +
                    user.username +
                    " )"}
                </p>
                <p>{user.email}</p>
              </div>
            </div>

            <div
              style={{
                display: "flex",
                flexDirection: "row",
                justifyContent: "space-between",
                width: "90%",
              }}
            >
              <Menu />
              <form className="form">
              <div><label htmlFor="startDate">start date</label></div>
              <div>
                <DatePicker
          controls={['calendar', 'time']}
        select="range"
        display="inline"
        touchUi={true}
        />
            </div>
            </form>
              <WaterLog />
            </div>
          </>
        )}
      </div>
    </div>
  );
};

export default WaterS;


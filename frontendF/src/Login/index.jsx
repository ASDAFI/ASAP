import { useState, useEffect } from "react";
import Cookies from "universal-cookie";
import { useNavigate } from "react-router-dom";

import { FInput } from "./FInput";

import "../App.css";

const Login = () => {
  const [fUsername, setUsername] = useState("");
  const [fPassword, setPassword] = useState("");
  const [buttonDisabled, setButtonDisabled] = useState(false);

  const cookies = new Cookies();
  const navigate = useNavigate();

  useEffect(() => {
    if (cookies.get("token")) {
      navigate("/main");
    }
  }, []);
  const handleSubmit = (username, password) => {
    setButtonDisabled(true);
    fetch("http://asap.carriot.ir:8000/login", {
      method: "POST",
      headers: {
        "content-type": "application/json",
      },
      body: JSON.stringify({
        username: username.toLowerCase(),
        password: password,
      }),
    }).then((response) => {
      if (response.status !== 200) {
        if (response.status === 500) {
          alert("user name or password is wrong\nplease try again");
        } else {
          alert("sorry an error occurred\nERROR CODE: " + response.status);
        }
        setButtonDisabled(false);
      } else {
        //alert("hmm");
        response.json().then((res) => {
          cookies.set("token", res.token, { path: "/" });
          navigate("/main");
        });
      }
    });
  };

  return (
    <div>
      <div
        style={{
          backgroundColor: "#ffffff94",
          borderRadius: 14,
          padding: 14,
          width: "40vw",
          textAlign: "center",
          boxShadow: "0px 0px 4px 2px #69fa97",
          marginTop: 44,
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <h1 className="Title">Farm Manager !</h1>
        <br style={{ width: "80%" }} />
        <h2>Login</h2>

        <form
          style={{ display: "flex", flexDirection: "column", width: "70%" }}
        >
          <FLabel text="username : " />
          <FInput setText={setUsername} />
          <FLabel text="password : " />
          <FInput setText={setPassword} />
          <button
            title="login"
            className="FButton"
            type="button"
            disabled={buttonDisabled}
            onClick={() => handleSubmit(fUsername, fPassword)}
          >
            login
          </button>
        </form>
      </div>

      <img
        src={require("../Assets/loginBG.jpg")}
        alt="bg"
        style={{
          width: "100%",
          position: "fixed",
          left: 0,
          top: 0,
          overflow: "hidden",
          zIndex: -1,
        }}
      />
    </div>
  );
};

const FLabel = ({ text }) => {
  return (
    <label
      style={{
        fontSize: 18,
        fontFamily: "Verdana",
        textAlign: "left",
        margin: 4,
        marginLeft: 14,
        marginBottom: 8,
      }}
    >
      {text}
    </label>
  );
};

export default Login;

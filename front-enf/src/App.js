import React from "react";
import NewLogin from "./component/form/NewLogin";
import Home from "./component/home/Home";
import "./index.css";

export const AuthContext = React.createContext();

const initialState = {
  isAuthenticated: false,
  user: null,
  token: null,
  error: null,
};
const reducer = (state, action) => {
  switch (action.type) {
    case "LOGIN":
      localStorage.setItem("user", JSON.stringify(action.payload.user));
      localStorage.setItem("token", JSON.stringify(action.payload.token));
      return {
        ...state,
        isAuthenticated: true,
        user: action.payload.user,
        token: action.payload.token,
        error: action.payload.error,
      };
    case "LOGOUT":
      localStorage.clear();
      return {
        ...state,
        isAuthenticated: false,
        user: null,
        error: action.payload.error,
      };
    default:
      return state;
  }
};

function App() {
  const onLoginHandler = async (data) => {
    console.log("data", data);
    try {
      let res = await fetch("http://localhost:8080/login/", {
        method: "POST",
        body: JSON.stringify({
          User: data.enteredUserName,
          Password: data.enteredUserPassword,
        }),
      });
      console.log("res", res);
      if (res.status === 200) {
        let resJson = await res.json();
        console.log("200 res", resJson);
        dispatch({
          type: "LOGIN",
          payload: {
            user: data.enteredUserName,
            error: null,
          },
        });
      }
      if (res.status === 401) {
        console.log("200 res", res.body);
        dispatch({
          type: "LOGOUT",
          payload: {
            user: data.enteredUserName,
            error: "Incorrect user or password",
          },
        });
      }
    } catch (err) {
      console.log("catch", err.statusText);
      console.log(err);
      dispatch({
        type: "LOGOUT",
        payload: {
          user: data.enteredUserName,
          error: "login error:" + err.message,
        },
      });
    }
  };

  const [state, dispatch] = React.useReducer(reducer, initialState);

  return (
    <div>
      <h3 className="error">{state.error}</h3>
      <div className="App">
        {!state.isAuthenticated ? (
          <NewLogin onLoginHandler={onLoginHandler} />
        ) : (
          <Home />
        )}
      </div>
    </div>
  );
}

export default App;

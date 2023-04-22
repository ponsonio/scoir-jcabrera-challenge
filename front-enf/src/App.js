import React from "react";
import NewLogin from "./component/form/NewLogin";
import Home from "./component/home/Home";

export const AuthContext = React.createContext();

const initialState = {
  isAuthenticated: false,
  user: null,
  token: null,
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
      };
    case "LOGOUT":
      localStorage.clear();
      return {
        ...state,
        isAuthenticated: false,
        user: null,
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
          },
        });
      }
      if (res.status === 401) {
        console.log("200 res", res.body);
      }
    } catch (err) {
      console.log("catch", err.statusText);
      console.log(err);
    }
  };

  const [state, dispatch] = React.useReducer(reducer, initialState);

  return (
    <div>
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

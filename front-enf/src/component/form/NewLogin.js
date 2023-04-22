import React from "react";
import "./NewLogin.css";
import LoginForm from "./LoginForm";
import Card from "../Card";

const NewLogin = (props) => {
  return (
    <Card className="expenses">
      <div className="new-expense">
        <LoginForm onLoginHandler={props.onLoginHandler} />
      </div>
    </Card>
  );
};

export default NewLogin;

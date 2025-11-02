import React from "react";
import styles from "../assets/css/Header.module.css";


export const Header: React.FC = () => (
    //アロー関数でアローの右側を（）で囲む場合はreturn不要
    <header className={styles.header}>
        <h1 className={styles.title}>ParkingHub</h1>
    </header>
);
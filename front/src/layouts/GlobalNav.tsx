import React from "react";
import styles from "../assets/css/Global.module.css";

const GlobalNav: React.FC = () => {
    return (
        <nav className={styles.navBar}>
            <ul className={styles.navList}>
                <li className={styles.navItem}><a href="/" className={styles.navLink}>HOME</a></li>
                <li className={styles.navItem}><a href="/owner" className={styles.navLink}>車両所有者</a></li>
            </ul>
        </nav>
    );
};

export default GlobalNav;
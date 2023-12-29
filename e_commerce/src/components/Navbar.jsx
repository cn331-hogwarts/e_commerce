import React from "react";
import { Link, NavLink } from 'react-router-dom';

function Navbar() {
    return (
        <>
            <nav className='navbar' style={{ background: `linear-gradient(90deg, #000000 10%)` }}>
                <div className="navbar-container">
                    <Link to='/' className='navbar-logo'>
                        nana
                    </Link>
                </div>
            </nav>
        </>
    );
}

export default Navbar;

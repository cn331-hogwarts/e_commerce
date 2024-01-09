import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Home from './components/pages/Home';
import Blog from './components/pages/Blog';
import Contact from './components/pages/Contact';
import Faq from './components/pages/Faq';
import Menus from './components/pages/Menus';
import Login from './components/pages/login';

function App() {
    return (
        <>
            <Router>
                <Navbar />
                <Routes>
                    <Route path='home' element={<Home />} />
                    <Route path='blog' element={<Blog />} />
                    <Route path='contact' element={<Contact />} />
                    <Route path='faq' element={<Faq />} />
                    <Route path='menus' element={<Menus />} />
                    <Route path='login' element={<Login />} />
                </Routes>
            </Router>
        </>
    );
}

export default App;

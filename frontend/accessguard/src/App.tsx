import { BrowserRouter as Router, Routes, Route, Navigate, useLocation } from "react-router-dom";


import './App.css'
import AdminDashboard from "./components/pages/AdminDashboard/AdminDashboard";
import SecurityEvents from "./components/pages/AdminDashboard/SecurityEvents/SecurityEvents";
import Login from "./components/pages/Login/Login";
import Users from "./components/pages/AdminDashboard/Users/Users";
import Roles from "./components/pages/AdminDashboard/Roles/Roles";
import Rules from "./components/pages/AdminDashboard/Rules/Rules";
import AuditLogs from "./components/pages/AdminDashboard/AuditLogs/AuditLogs";
import LoginHistory from "./components/pages/AdminDashboard/LoginHistory/LoginHistory";
import FileDashboard from "./components/pages/FilesDashboard/FilesDashboard";
import Navbar from "./components/navbar/nav";

function App() {
 
  const showNavbar =
    location.pathname !== "/login" ;

  return (
    <>
      <Router>
        {showNavbar && <Navbar/>}
        <div>
          <Routes>

            <Route path="/" element={<Navigate to="/Login" />} />
            <Route path="/Login" element={<Login />} />
            <Route path="/files" element={<FileDashboard />} />
            <Route path="/admin" element={<AdminDashboard />}>
              <Route path="users" element={<Users />} />
              <Route path="roles" element={<Roles />} />
              <Route path="rules" element={<Rules />} />
              <Route path="audit-logs" element={<AuditLogs />} />
              <Route path="login-history" element={<LoginHistory />} />
              <Route path="security-events" element={<SecurityEvents />} />
            </Route>
          </Routes>
        </div>
      </Router>
    </>
  )
}

export default App

import { Navigate } from 'react-router-dom';
import { AuthContext } from '../context/AuthContext';
import { useContext } from 'react';

function PublicRoute({children}) {
    const { isAuthenticated } = useContext(AuthContext);
    if(isAuthenticated==null) return null;
    if(isAuthenticated){
        return <Navigate to="/home" replace/>
    }
    return children;
}

export default PublicRoute

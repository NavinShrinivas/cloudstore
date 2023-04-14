import { useState } from 'react'
import { useSelector } from 'react-redux'

import { HomeNavbar } from '../components'

function SellerHomePage() {


    const [user, setUser] = useState(useSelector((state) => state.user.value))

    return (
        <>
            <div className="home-page" style={{ backgroundColor: '#f5f5f5', height: '100vh', overflow: 'hidden' }}>
                <HomeNavbar />
                <div className="mx-4" style={{ marginTop: '70px' }}>

                </div>
            </div>
        </>
    );
}

export default SellerHomePage;

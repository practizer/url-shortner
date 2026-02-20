import React from 'react'
import { GoogleLogin } from "@react-oauth/google";
import { useNavigate } from "react-router-dom";
import api from '../api/axios';

function Login() {
    const navigate = useNavigate();

    const handleSuccess = async (credentialResponse) => {
        try {
            const res = await api.post("/auth/login", {
                token: credentialResponse.credential,
            });
            if (res.status === 200) {
                navigate("/home");
            }
        } catch (error) {
            alert("Login Failed");
            console.log(error);
        }
    };

    return (
        <div className="min-h-screen w-screen bg-zinc-950 flex items-center justify-center relative overflow-hidden">
            
            <div
                className="absolute inset-0 opacity-[0.04]"
                style={{ backgroundImage: `linear-gradient(#fff 1px, transparent 1px), linear-gradient(90deg, #fff 1px, transparent 1px)`,backgroundSize: '48px 48px',}}
            />
            
            <div className="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[500px] h-[500px] bg-indigo-600/10 rounded-full blur-[120px] pointer-events-none" />

            <div className="relative z-10 w-full max-w-sm mx-4">
                <div className="bg-zinc-900 border border-zinc-800 rounded-2xl shadow-2xl shadow-black/60 overflow-hidden">

                    <div className="h-0.5 w-full bg-gradient-to-r from-indigo-500 via-violet-500 to-indigo-500" />

                    <div className="px-10 py-12 flex flex-col items-center gap-8">
                        
                        <div className="flex flex-col items-center gap-2 text-center">
                            <div className="flex items-center justify-center w-14 h-14 rounded-xl bg-indigo-600/15 border border-indigo-500/30 mb-1">
                                <svg className="w-7 h-7 text-indigo-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.8}>
                                    <path strokeLinecap="round" strokeLinejoin="round" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
                                </svg>
                            </div>
                            <h1 className="text-white text-3xl font-bold tracking-tight" style={{ fontFamily: "'Georgia', serif", letterSpacing: '-0.02em' }}>
                                BITS Links
                            </h1>
                            <p className="text-zinc-400 text-sm tracking-widest uppercase font-medium">
                                The URL Shortener
                            </p>
                        </div>
                        
                        <div className="w-full flex items-center gap-3">
                            <div className="flex-1 h-px bg-zinc-800" />
                            <span className="text-zinc-600 text-xs tracking-widest uppercase">Sign in</span>
                            <div className="flex-1 h-px bg-zinc-800" />
                        </div>

                        <div className="flex flex-col items-center gap-4 w-full">
                            <div className="w-full flex justify-center [&>div]:w-full [&>div>div]:w-full [&>div>div>div]:w-full [&_iframe]:w-full">
                                <GoogleLogin
                                    onSuccess={handleSuccess}
                                    onError={() => alert("Google Login Failed")}
                                    theme="filled_black"
                                    size="large"
                                    shape="rectangular"
                                    width="100%"
                                />
                            </div>
                            <p className="text-zinc-600 text-xs text-center leading-relaxed">
                                By signing in, you agree to our{' '}
                                <span className="text-zinc-400 underline underline-offset-2 cursor-pointer hover:text-white transition-colors">Terms</span>
                                {' '}and{' '}
                                <span className="text-zinc-400 underline underline-offset-2 cursor-pointer hover:text-white transition-colors">Privacy Policy</span>.
                            </p>
                        </div>

                    </div>
                </div>
                
                <p className="text-center text-zinc-700 text-xs mt-5 tracking-wide">
                    © 2026 BITS · Shorten smarter.
                </p>
            </div>
        </div>
    );
}

export default Login;
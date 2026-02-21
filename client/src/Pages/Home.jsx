import React, { useState } from 'react'
import NavBar from '../components/NavBar'
import api from '../api/axios'

const BASE_URL = 'https://localhost:5000'

function Home() {
    const [originalUrl, setOriginalUrl] = useState('')
    const [shortCode, setShortCode] = useState('')
    const [availability, setAvailability] = useState(null)
    const [submitStatus, setSubmitStatus] = useState(null)
    const [submitMessage, setSubmitMessage] = useState('')
    const [loading, setLoading] = useState(false)
    const [createdUrl, setCreatedUrl] = useState('')
    const [copied, setCopied] = useState(false)

    const sleep = (ms) => new Promise((resolve) => setTimeout(resolve, ms))

    const checkAvailability = async () => {
        if (!shortCode.trim()) return
        setAvailability('checking')
        try {
            await sleep(800)
            const res = await api.post('/check', { short_code: shortCode })
            setAvailability(res.data.status ? 'available' : 'taken')
        } catch (err) {
            setAvailability('taken')
        }
    }

    const handleShortCodeChange = (e) => {
        setShortCode(e.target.value)
        setAvailability(null)
        setSubmitStatus(null)
        setCreatedUrl('')
    }

    const handleSubmit = async () => {
        if (!originalUrl.trim() || !shortCode.trim()) return
        if (availability !== 'available') return
        setLoading(true)
        setSubmitStatus(null)
        setCreatedUrl('')
        try {
            const res = await api.post('/url/add', {
                short_code: shortCode,
                original_url: originalUrl,
            })
            setSubmitStatus('success')
            setSubmitMessage(res.data.message || 'URL created successfully!')
            setCreatedUrl(`${BASE_URL}/${shortCode}`)
            setOriginalUrl('')
            setShortCode('')
            setAvailability(null)
        } catch (err) {
            setSubmitStatus('error')
            setSubmitMessage(err?.response?.data?.error || 'Something went wrong.')
        } finally {
            setLoading(false)
        }
    }

    const handleCopy = () => {
        navigator.clipboard.writeText(createdUrl).then(() => {
            setCopied(true)
            setTimeout(() => setCopied(false), 2000)
        })
    }

    const availabilityColor = {
        available: 'text-emerald-400',
        taken: 'text-red-400',
        checking: 'text-zinc-400',
    }

    const availabilityText = {
        available: '✓ Short code is available',
        taken: '✗ Short code is already taken',
        checking: 'Checking availability...',
    }

    return (
        <div className="min-h-screen w-screen bg-zinc-950 flex flex-col relative overflow-hidden">

            {/* Grid background */}
            <div
                className="absolute inset-0 opacity-[0.04] pointer-events-none"
                style={{
                    backgroundImage: `linear-gradient(#fff 1px, transparent 1px), linear-gradient(90deg, #fff 1px, transparent 1px)`,
                    backgroundSize: '48px 48px',
                }}
            />

            <div className="absolute top-1/3 left-1/2 -translate-x-1/2 -translate-y-1/2 w-[300px] h-[300px] sm:w-[600px] sm:h-[600px] bg-indigo-600/10 rounded-full blur-[140px] pointer-events-none" />

            <NavBar />

            <main className="relative z-10 flex-1 flex items-center justify-center px-4 py-12">
                <div className="w-full max-w-lg flex flex-col gap-6">

                    {/* Header */}
                    <div className="text-center mb-2">
                        <h2
                            className="text-white text-2xl font-bold tracking-tight"
                            style={{ fontFamily: "'Georgia', serif", letterSpacing: '-0.02em' }}
                        >
                            Shorten a URL
                        </h2>
                        <p className="text-zinc-500 text-sm mt-1">Create a short, memorable link in seconds.</p>
                    </div>

                    {/* Card */}
                    <div className="bg-zinc-900 border border-zinc-800 rounded-2xl shadow-2xl shadow-black/60 overflow-hidden">

                        <div className="h-0.5 w-full bg-gradient-to-r from-indigo-500 via-violet-500 to-indigo-500" />

                        <div className="px-4 sm:px-8 py-6 sm:py-8 flex flex-col gap-6">

                            {/* Original URL */}
                            <div className="flex flex-col gap-2">
                                <label className="text-zinc-400 text-xs font-medium uppercase tracking-widest">
                                    Destination URL
                                </label>
                                <input
                                    type="url"
                                    placeholder="https://example.com/long/url"
                                    value={originalUrl}
                                    onChange={(e) => {
                                        setOriginalUrl(e.target.value)
                                        setSubmitStatus(null)
                                        setCreatedUrl('')
                                    }}
                                    className="w-full bg-zinc-950 border border-zinc-800 rounded-xl px-4 py-3 text-white text-sm placeholder-zinc-600 focus:outline-none focus:border-indigo-500/60 focus:ring-1 focus:ring-indigo-500/30 transition-all"
                                />
                            </div>

                            {/* Divider */}
                            <div className="flex items-center gap-3">
                                <div className="flex-1 h-px bg-zinc-800" />
                                <span className="text-zinc-600 text-xs tracking-widest uppercase">Short code</span>
                                <div className="flex-1 h-px bg-zinc-800" />
                            </div>

                            {/* Short code + check */}
                            <div className="flex flex-col gap-2">
                                <label className="text-zinc-400 text-xs font-medium uppercase tracking-widest">
                                    Custom Short Code
                                </label>

                                {/* Base URL hint — visible on mobile above the input */}
                                <p className="text-zinc-600 text-xs sm:hidden break-all">
                                    {BASE_URL}/
                                </p>

                                <div className="flex gap-2 items-stretch">
                                    {/* Input box */}
                                    <div className="flex items-center bg-zinc-950 border border-zinc-800 rounded-xl px-3 sm:px-4 py-3 text-sm flex-1 min-w-0 gap-1 focus-within:border-indigo-500/60 focus-within:ring-1 focus-within:ring-indigo-500/30 transition-all overflow-hidden">
                                        {/* Base URL prefix — hidden on mobile, shown on sm+ */}
                                        <span className="hidden sm:inline text-zinc-600 select-none whitespace-nowrap shrink-0 text-xs">
                                            {BASE_URL}/
                                        </span>
                                        <input
                                            type="text"
                                            placeholder="my-link"
                                            value={shortCode}
                                            onChange={handleShortCodeChange}
                                            className="bg-transparent text-white placeholder-zinc-600 focus:outline-none flex-1 min-w-0 text-sm"
                                        />
                                    </div>

                                    {/* Check button */}
                                    <button
                                        onClick={checkAvailability}
                                        disabled={!shortCode.trim() || availability === 'checking'}
                                        className="shrink-0 px-3 sm:px-4 py-3 rounded-xl bg-zinc-800 hover:bg-zinc-700 border border-zinc-700 text-zinc-300 text-sm font-medium transition-colors disabled:opacity-40 disabled:cursor-not-allowed whitespace-nowrap"
                                    >
                                        {availability === 'checking' ? (
                                            <span className="flex items-center gap-1.5">
                                                <svg className="animate-spin h-3.5 w-3.5 text-zinc-400" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24">
                                                    <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" />
                                                    <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8v8H4z" />
                                                </svg>
                                                <span className="hidden sm:inline">Checking</span>
                                            </span>
                                        ) : 'Check'}
                                    </button>
                                </div>

                                {/* Availability badge */}
                                {availability && (
                                    <p className={`text-xs mt-0.5 transition-all ${availabilityColor[availability]}`}>
                                        {availabilityText[availability]}
                                    </p>
                                )}
                            </div>

                            {/* Submit */}
                            <button
                                onClick={handleSubmit}
                                disabled={loading || !originalUrl.trim() || availability !== 'available'}
                                className="w-full py-3 rounded-xl bg-indigo-600 hover:bg-indigo-500 disabled:opacity-40 disabled:cursor-not-allowed text-white text-sm font-semibold tracking-wide transition-colors shadow-lg shadow-indigo-900/30"
                            >
                                {loading ? 'Adding...' : 'Create Short Link'}
                            </button>

                            {/* Success feedback */}
                            {submitStatus === 'success' && (
                                <div className="rounded-xl border border-emerald-500/30 bg-emerald-500/10 overflow-hidden">
                                    <div className="px-4 py-3 text-sm text-center text-emerald-400">
                                        {submitMessage}
                                    </div>
                                    {createdUrl && (
                                        <>
                                            <div className="h-px bg-emerald-500/20" />
                                            <div className="px-4 py-3 flex flex-wrap items-center gap-2">
                                                <span className="text-emerald-300 text-sm font-mono min-w-0 flex-1 truncate break-all">
                                                    {createdUrl}
                                                </span>
                                                <div className="flex items-center gap-2 shrink-0">
                                                    <button
                                                        onClick={handleCopy}
                                                        className="px-3 py-1.5 rounded-lg bg-emerald-500/20 hover:bg-emerald-500/30 border border-emerald-500/30 text-emerald-400 text-xs font-medium transition-colors whitespace-nowrap"
                                                    >
                                                        {copied ? '✓ Copied!' : 'Copy'}
                                                    </button>
                                                    <a
                                                        href={createdUrl}
                                                        target="_blank"
                                                        rel="noopener noreferrer"
                                                        className="px-3 py-1.5 rounded-lg bg-emerald-500/20 hover:bg-emerald-500/30 border border-emerald-500/30 text-emerald-400 text-xs font-medium transition-colors whitespace-nowrap"
                                                    >
                                                        Open ↗
                                                    </a>
                                                </div>
                                            </div>
                                        </>
                                    )}
                                </div>
                            )}

                            {/* Error feedback */}
                            {submitStatus === 'error' && (
                                <div className="rounded-xl px-4 py-3 text-sm text-center border bg-red-500/10 border-red-500/30 text-red-400">
                                    {submitMessage}
                                </div>
                            )}

                        </div>
                    </div>

                    <p className="text-center text-zinc-700 text-xs tracking-wide">
                        © 2026 BITS · Shorten smarter.
                    </p>
                </div>
            </main>
        </div>
    )
}

export default Home
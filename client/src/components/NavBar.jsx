import React from 'react'

function NavBar() {
  return (
    <div className="w-full h-16 bg-zinc-900 border-b border-zinc-800 flex items-center px-8 relative">
      <div className="h-0.5 w-full bg-gradient-to-r from-indigo-500 via-violet-500 to-indigo-500 absolute top-0 left-0" />
      <div className="flex items-center gap-3">
        <div className="flex items-center justify-center w-8 h-8 rounded-lg bg-indigo-600/15 border border-indigo-500/30">
          <svg className="w-4 h-4 text-indigo-400" fill="none" viewBox="0 0 24 24" stroke="currentColor" strokeWidth={1.8}>
            <path strokeLinecap="round" strokeLinejoin="round" d="M13.828 10.172a4 4 0 00-5.656 0l-4 4a4 4 0 105.656 5.656l1.102-1.101m-.758-4.899a4 4 0 005.656 0l4-4a4 4 0 00-5.656-5.656l-1.1 1.1" />
          </svg>
        </div>
        <span className="text-white text-xl font-bold tracking-tight" style={{ fontFamily: "'Georgia', serif", letterSpacing: '-0.02em' }}>
          BITS Links
        </span>
      </div>
    </div>
  )
}

export default NavBar
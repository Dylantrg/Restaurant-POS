import { useEffect, useState } from 'react'

function App() {
  const [health, setHealth] = useState('...')

  useEffect(() => {
    fetch('/api/health')
      .then(r => r.json())
      .then(d => setHealth(d.status))
      .catch(() => setHealth('error'))
  }, [])

  return (
    <div style={{ fontFamily: 'system-ui, -apple-system, Segoe UI, Roboto, sans-serif', padding: 24 }}>
      <h1>POS Management</h1>
      <p>Backend health: <strong>{health}</strong></p>
    </div>
  )
}

export default App

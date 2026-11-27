// Pruebas automáticas del programa. Uso:  cd tests && npm install && npx playwright install chromium && npm test
// Cada prueba abre el programa en Chromium (sin pantalla) y comprueba un comportamiento con assert.
const {chromium}=require('playwright');
const assert=require('assert/strict');
const fs=require('fs'),path=require('path'),os=require('os');
const APP='file://'+path.resolve(__dirname,'../app/DisenoElectrico.html');
const TMP=fs.mkdtempSync(path.join(os.tmpdir(),'de-pruebas-'));
const pruebas=[];const prueba=(n,f)=>pruebas.push({n,f});

async function abrir(browser,{sample=true}={}){
 const ctx=await browser.newContext({viewport:{width:1440,height:860},acceptDownloads:true});const p=await ctx.newPage();
 p.errores=[];p.on('pageerror',e=>p.errores.push(e.message));
 p.on('dialog',d=>d.accept(d.type()==='prompt'?'12':undefined));
 await p.route('https://api.github.com/**',r=>r.abort());
 await p.goto(APP);await p.waitForFunction(()=>typeof state==='object'&&typeof render==='function');
 await p.waitForTimeout(500); // la ayuda de bienvenida aparece a los 300 ms
 await p.evaluate(s=>{document.querySelector('#modal').classList.remove('show');if(tourI>=0)tourEnd();settings.autoSw=false;if(s){buildSample();state.pid=uid();changed()}},sample);
 return p}

// ---------------------------------------------------------------- simulación
prueba('Casa de ejemplo: abre sin errores y la revisión no tiene errores graves',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>({dev:state.devices.length,wires:state.wires.length,err:check().filter(i=>i.l==='error').map(i=>i.m)}));
 assert.ok(r.dev>40&&r.wires>50,'la casa de ejemplo debe tener aparatos y cables');assert.deepEqual(r.err,[]);return p});
prueba('Interruptor sencillo enciende y apaga el bombillo',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>{setMode('sim');const k=state.devices.find(d=>d.type==='int_sencillo'&&roomOf(d)&&roomOf(d).type==='cocina'),l=state.devices.find(d=>d.type==='bombillo'&&roomOf(d)&&roomOf(d).type==='cocina');
  const on=()=>{render();return !!(simRes.dev[l.id]&&simRes.dev[l.id].on)};k.st[0]=false;const a=on();k.st[0]=true;const c=on();return [a,c]});
 assert.deepEqual(r,[false,true]);return p});
prueba('Conmutables y cruce: el bombillo cambia con cada uno de los 3 interruptores',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>{setMode('sim');const l=state.devices.find(d=>d.type==='bombillo'&&roomOf(d).type==='alcoba');
  const sw=[...state.devices.filter(d=>d.type==='int_conmutable'&&roomOf(d).type==='alcoba'),state.devices.find(d=>d.type==='int_cruce')];
  const on=()=>{render();return !!(simRes.dev[l.id]&&simRes.dev[l.id].on)};const s=[on()];sw.forEach(x=>{x.st[0]=!x.st[0];s.push(on())});return {n:sw.length,s}});
 assert.equal(r.n,3);for(let i=1;i<r.s.length;i++) assert.notEqual(r.s[i],r.s[i-1],'cada interruptor debe cambiar el bombillo');return p});
prueba('Tomacorrientes energizados y con tierra',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>{setMode('sim');render();return state.devices.filter(d=>CAT[d.type].kind==='toma').map(d=>simRes.dev[d.id]).filter(x=>!x||!x.on||!x.gnd).length});
 assert.equal(r,0);return p});
prueba('Cortocircuito: dispara el taco al instante',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>{const t=state.devices.find(d=>d.type==='toma_doble');state.wires.push({id:'x',a:{d:t.id,t:'F'},b:{d:t.id,t:'N'},pts:[],color:'negro',gauge:'12'});
  const res=simulate();const c=state.circuits.filter(c=>c.tripped);return {msg:res.msgs.join(' '),n:c.length}});
 assert.equal(r.n,1);assert.match(r.msg,/Cortocircuito/);return p});
prueba('Falla a tierra en circuito con diferencial: dispara el diferencial',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>{const t=state.devices.find(d=>d.type==='toma_gfci');const c=circ(t.c);c.dif=true;
  state.wires.push({id:'x',a:{d:t.id,t:'F'},b:{d:t.id,t:'T'},pts:[],color:'negro',gauge:'12'});const res=simulate();return {msg:res.msgs.join(' '),trip:c.tripped}});
 assert.equal(r.trip,true);assert.match(r.msg,/diferencial/i);return p});
prueba('GFCI: TEST corta el tomacorriente y RESET lo restablece',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>{setMode('sim');const t=state.devices.find(d=>d.type==='toma_gfci');const on=()=>{render();return !!simRes.dev[t.id].on};
  const a=on();gfciTest(t);const c=on();gfciReset(t);return [a,c,on()]});
 assert.deepEqual(r,[true,false,true]);return p});
prueba('Sobrecarga: el taco se calienta y se dispara',async b=>{
 const p=await abrir(b);await p.evaluate(()=>{setMode('sim');const ts=state.devices.filter(d=>d.type==='toma_doble'&&roomOf(d)&&roomOf(d).type==='sala');ts.forEach(t=>t.plug=APPL.findIndex(a=>a&&a.w>=1500));changed()});
 await p.waitForFunction(()=>state.circuits.some(c=>c.tripped),null,{timeout:20000});return p});

// ---------------------------------------------------------------- diseño
prueba('Diseño automático de una casa de 4 alcobas: rápido y sin errores graves',async b=>{
 const p=await abrir(b,{sample:false});const r=await p.evaluate(()=>{state=newState();state.pid=uid();insertTemplate('casa4',true);RW=roomWalls();state.devices.push({id:'P',type:'tablero',x:150,y:120,r:0,st:[]});defaultCircuits();
  const t=performance.now();autoDesignAll('todo');const ms=performance.now()-t;render();return {ms,dev:state.devices.filter(isElec).length,wires:state.wires.length,err:check().filter(i=>i.l==='error').length}});
 assert.ok(r.dev>40&&r.wires>80,JSON.stringify(r));assert.ok(r.ms<3000,'demasiado lento: '+r.ms);assert.equal(r.err,0);return p});
prueba('Deshacer y rehacer',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>{const n=state.devices.length;pushUndo();state.devices.push({id:'zz',type:'bombillo',x:10,y:10,r:0});changed();undo();const a=state.devices.length;redo();return [n,a,state.devices.length]});
 assert.deepEqual(r,[r[0],r[0],r[0]+1]);return p});
prueba('Tablero: cambiar el número de espacios se refleja en el plano',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>{const P=getPanel();const h0=document.querySelector('#layer').innerHTML.length;PNL().slots=24;changed();return [h0!==document.querySelector('#layer').innerHTML.length,PNL().slots]});
 assert.deepEqual(r,[true,24]);return p});
prueba('Cuadro de cargas, materiales, presupuesto, unifilar e impresión se generan',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>{const L=loadTable(),M=materials(),R=budgetRows(),U=unifilarSVG();return {l:L.rows?L.rows.length:Object.keys(L).length,m:Object.keys(M.dev).length,r:R.length,u:U.length}});
 assert.ok(r.l>0&&r.m>0&&r.r>0&&r.u>1000,JSON.stringify(r));
 const [pop]=await Promise.all([p.waitForEvent('popup',{timeout:10000}).catch(()=>null),p.evaluate(()=>printDoc())]);
 if(pop){await pop.waitForLoadState();assert.ok((await pop.content()).length>5000)}return p});

// ---------------------------------------------------------------- archivos y seguridad
prueba('Guardar y volver a abrir un archivo conserva el proyecto',async b=>{
 const p=await abrir(b);await p.evaluate(()=>window.showSaveFilePicker=undefined);
 const [dl]=await Promise.all([p.waitForEvent('download'),p.evaluate(()=>saveFile(false))]);const f=path.join(TMP,'p.elec.json');await dl.saveAs(f);
 const before=await p.evaluate(()=>[state.devices.length,state.wires.length,state.circuits.length]);
 await p.evaluate(()=>{newProject('vacío')});await p.setInputFiles('#fileIn',f);await p.waitForTimeout(600);
 assert.deepEqual(await p.evaluate(()=>[state.devices.length,state.wires.length,state.circuits.length]),before);return p});
prueba('Archivo malicioso: no ejecuta código',async b=>{
 const p=await abrir(b);const evil=await p.evaluate(()=>{const s=JSON.parse(ser());const X='"><img src=x onerror="window.__pw=(window.__pw||0)+1">';
  s.name=X;s.rooms[0].name=X;s.circuits[0].name=X;s.devices[3].label=X;s.devices[4].x='0" onmouseover="window.__pw++';s.colors.retorno='azul" onload="window.__pw++';
  s.bg={src:'x" onload="window.__pw++" y="',x:0,y:0,w:100,h:100,op:.5};s.wires[0].c=X;s.info={dir:X};return JSON.stringify(s)});
 const f=path.join(TMP,'evil.json');fs.writeFileSync(f,evil);await p.setInputFiles('#fileIn',f);await p.waitForTimeout(400);
 await p.evaluate(()=>{['prop','panel','check','mat','cost'].forEach(t=>{sideTab=t;try{renderSide()}catch(e){}})});await p.waitForTimeout(600);
 assert.equal(await p.evaluate(()=>window.__pw||0),0);return p});
prueba('Archivo dañado: no se abre y el proyecto actual queda intacto',async b=>{
 const p=await abrir(b);const n=await p.evaluate(()=>state.devices.length);
 for(const [k,txt] of [['a','{"devices":['],['b','{"devices":[{"type":"bombillo"}],"rooms":"x"}'],['c','[1,2,3]'],['d','{"devices":"x"}']]){
  const f=path.join(TMP,k+'.json');fs.writeFileSync(f,txt);await p.setInputFiles('#fileIn',f);await p.waitForTimeout(300);
  assert.equal(await p.evaluate(()=>state.devices.length),n,'archivo '+k);assert.match(await p.textContent('#toast'),/No se abrió/)}
 await p.evaluate(()=>render());return p});
prueba('Imagen de fondo grande: se reduce y se guarda aparte del proyecto',async b=>{
 const p=await abrir(b);const png=await p.evaluate(()=>{const c=document.createElement('canvas');c.width=5000;c.height=3500;const g=c.getContext('2d');for(let i=0;i<4000;i++){g.fillStyle='hsl('+i+',70%,50%)';g.fillRect(Math.random()*5000,Math.random()*3500,50,50)}return c.toDataURL('image/png')});
 const f=path.join(TMP,'plano.png');fs.writeFileSync(f,Buffer.from(png.split(',')[1],'base64'));await p.setInputFiles('#imgIn',f);
 await p.waitForFunction(()=>state.bg&&state.bg.img,null,{timeout:10000});
 const r=await p.evaluate(()=>({kb:JSON.stringify(state).length/1024,img:!!document.querySelector('#layer image'),w:BGI[state.bg.img].length}));
 assert.ok(r.img&&r.kb<200,'el proyecto no debe incluir la imagen: '+r.kb);
 await p.reload();await p.waitForFunction(()=>document.querySelector('#layer image'),null,{timeout:10000});return p});
prueba('Respaldo completo y restauración en un navegador limpio',async b=>{
 const p=await abrir(b);await p.evaluate(()=>{state.name='Proyecto respaldado';changed();openLib()});
 const [dl]=await Promise.all([p.waitForEvent('download',{timeout:8000}),p.click('#lbBk')]).catch(async e=>{throw new Error('toast='+await p.textContent('#toast')+' errs='+p.errores)});const f=path.join(TMP,'respaldo.json');await dl.saveAs(f);
 const q=await abrir(b,{sample:false});await q.setInputFiles('#bkIn',f);await q.waitForFunction(()=>/restaurado/.test(document.querySelector('#toast').textContent),null,{timeout:10000});
 assert.ok(await q.evaluate(()=>libIndex().some(x=>x.name==='Proyecto respaldado')));await q.context().close();return p});
prueba('Versiones del proyecto: guardar y restaurar',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>{sideTab='prop';sel.clear();renderSide();const n=state.devices.length;$('#snapNew').click();state.devices.splice(0,5);changed();renderSide();document.querySelector('[data-sr]').click();return [n,state.devices.length]});
 assert.equal(r[0],r[1]);return p});

prueba('Recorrido de bienvenida: aparece la primera vez y recorre los 8 pasos',async b=>{
 const ctx=await b.newContext({viewport:{width:1366,height:768}});const p=await ctx.newPage();p.errores=[];p.on('pageerror',e=>p.errores.push(e.message));await p.route('https://api.github.com/**',r=>r.abort());
 await p.goto(APP);await p.waitForSelector('#tourBox',{timeout:5000});let n=1;
 while(await p.$('#tourNext')){const t=await p.textContent('#tourNext');await p.click('#tourNext');if(t==='Terminar')break;n++}
 assert.equal(n,8);assert.equal(await p.$('#tourBox'),null);
 await p.reload();await p.waitForTimeout(800);assert.equal(await p.$('#tourBox'),null,'no debe repetirse');return p});

// ---------------------------------------------------------------- 3D
prueba('Vista 3D: se construye y dibuja',async b=>{
 const p=await abrir(b);const r=await p.evaluate(()=>{v3Open();return new Promise(ok=>setTimeout(()=>ok({open:!!V3&&!!V3.scene,n:V3.scene?V3.scene.children.length:0}),800))});
 assert.ok(r.open&&r.n>5,JSON.stringify(r));await p.evaluate(()=>v3Close());return p});

// ---------------------------------------------------------------- ejecutar
(async()=>{
 const b=await chromium.launch({args:['--use-gl=angle','--use-angle=swiftshader','--enable-unsafe-swiftshader']});
 let ok=0,mal=0;const solo=process.argv[2];
 for(const t of pruebas){if(solo&&!t.n.toLowerCase().includes(solo.toLowerCase()))continue;const t0=Date.now();let p;
  try{p=await t.f(b);if(p&&p.errores.length)throw new Error('errores de JavaScript: '+p.errores.join(' | '));ok++;console.log(`✔ ${t.n} (${Date.now()-t0} ms)`)}
  catch(e){mal++;console.log(`✘ ${t.n}\n   ${String(e.message).split('\n').slice(0,6).join('\n   ')}`)}
  finally{if(p)await p.context().close().catch(()=>{})}}
 await b.close();console.log(`\n${ok} pruebas correctas, ${mal} fallidas`);process.exit(mal?1:0)})();

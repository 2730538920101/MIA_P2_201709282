// Generated from /home/ec2-user/MIA_P2/Cmd.g4 by ANTLR 4.8

        import "../ast"
        import "../ast/parametros"
        import arrayList "github.com/colegno/arraylist"
        import "strings"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class CmdParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		TOK_IGUAL=1, TOK_MKDISK=2, TOK_RMDISK=3, TOK_FDISK=4, TOK_MOUNT=5, TOK_MKFS=6, 
		TOK_LOGIN=7, TOK_LOGOUT=8, TOK_MKGRP=9, TOK_RMGRP=10, TOK_MKUSR=11, TOK_RMUSR=12, 
		TOK_MKFILE=13, TOK_MKDIR=14, TOK_EXEC=15, TOK_PAUSE=16, TOK_REP=17, TOK_PATH=18, 
		TOK_FIT=19, TOK_SIZE=20, TOK_UNIT=21, TOK_NAME=22, TOK_TYPE=23, TOK_ID=24, 
		TOK_USUARIO=25, TOK_PASSWORD=26, TOK_PWD=27, TOK_CONT=28, TOK_GRP=29, 
		TOK_RUTA=30, TOK_R=31, TOK_P=32, TOK_KB=33, TOK_MB=34, TOK_BYTES=35, TOK_PRIMARIA=36, 
		TOK_EXTENDIDA=37, TOK_LOGICA=38, TOK_FIRST=39, TOK_WORST=40, TOK_BEST=41, 
		TOK_FAST=42, TOK_FULL=43, TOK_CADENA=44, TOK_NUMERO=45, TOK_IDENTIFICADOR=46, 
		TOK_CAMINO=47, TOK_PALABRA=48, COMENTARIOS=49, WHITESPACE=50;
	public static final int
		RULE_start = 0, RULE_comandList = 1, RULE_comando = 2, RULE_param_list = 3, 
		RULE_param = 4;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "comandList", "comando", "param_list", "param"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "TOK_IGUAL", "TOK_MKDISK", "TOK_RMDISK", "TOK_FDISK", "TOK_MOUNT", 
			"TOK_MKFS", "TOK_LOGIN", "TOK_LOGOUT", "TOK_MKGRP", "TOK_RMGRP", "TOK_MKUSR", 
			"TOK_RMUSR", "TOK_MKFILE", "TOK_MKDIR", "TOK_EXEC", "TOK_PAUSE", "TOK_REP", 
			"TOK_PATH", "TOK_FIT", "TOK_SIZE", "TOK_UNIT", "TOK_NAME", "TOK_TYPE", 
			"TOK_ID", "TOK_USUARIO", "TOK_PASSWORD", "TOK_PWD", "TOK_CONT", "TOK_GRP", 
			"TOK_RUTA", "TOK_R", "TOK_P", "TOK_KB", "TOK_MB", "TOK_BYTES", "TOK_PRIMARIA", 
			"TOK_EXTENDIDA", "TOK_LOGICA", "TOK_FIRST", "TOK_WORST", "TOK_BEST", 
			"TOK_FAST", "TOK_FULL", "TOK_CADENA", "TOK_NUMERO", "TOK_IDENTIFICADOR", 
			"TOK_CAMINO", "TOK_PALABRA", "COMENTARIOS", "WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Cmd.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public CmdParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public ast.Ast ast;
		public ComandListContext comandList;
		public ComandListContext comandList() {
			return getRuleContext(ComandListContext.class,0);
		}
		public TerminalNode EOF() { return getToken(CmdParser.EOF, 0); }
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(10);
			((StartContext)_localctx).comandList = comandList(0);
			_localctx.ast = ast.NewAst(((StartContext)_localctx).comandList.lista)
			setState(12);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ComandListContext extends ParserRuleContext {
		public *arrayList.List lista;
		public ComandListContext AUXCLIST;
		public ComandoContext comando;
		public ComandoContext comando() {
			return getRuleContext(ComandoContext.class,0);
		}
		public ComandListContext comandList() {
			return getRuleContext(ComandListContext.class,0);
		}
		public ComandListContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comandList; }
	}

	public final ComandListContext comandList() throws RecognitionException {
		return comandList(0);
	}

	private ComandListContext comandList(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ComandListContext _localctx = new ComandListContext(_ctx, _parentState);
		ComandListContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_comandList, _p);
		 _localctx.lista = arrayList.New()
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(15);
			((ComandListContext)_localctx).comando = comando();

			                        _localctx.lista.Add(((ComandListContext)_localctx).comando.com)
			                   
			}
			_ctx.stop = _input.LT(-1);
			setState(24);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ComandListContext(_parentctx, _parentState);
					_localctx.AUXCLIST = _prevctx;
					_localctx.AUXCLIST = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_comandList);
					setState(18);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(19);
					((ComandListContext)_localctx).comando = comando();

					                                  ((ComandListContext)_localctx).AUXCLIST.lista.Add(((ComandListContext)_localctx).comando.com)
					                                  _localctx.lista = ((ComandListContext)_localctx).AUXCLIST.lista
					                             
					}
					} 
				}
				setState(26);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ComandoContext extends ParserRuleContext {
		public ast.Command com;
		public Param_listContext param_list;
		public TerminalNode TOK_MKDISK() { return getToken(CmdParser.TOK_MKDISK, 0); }
		public Param_listContext param_list() {
			return getRuleContext(Param_listContext.class,0);
		}
		public TerminalNode TOK_RMDISK() { return getToken(CmdParser.TOK_RMDISK, 0); }
		public TerminalNode TOK_FDISK() { return getToken(CmdParser.TOK_FDISK, 0); }
		public TerminalNode TOK_MOUNT() { return getToken(CmdParser.TOK_MOUNT, 0); }
		public TerminalNode TOK_MKFS() { return getToken(CmdParser.TOK_MKFS, 0); }
		public TerminalNode TOK_LOGIN() { return getToken(CmdParser.TOK_LOGIN, 0); }
		public TerminalNode TOK_MKGRP() { return getToken(CmdParser.TOK_MKGRP, 0); }
		public TerminalNode TOK_RMGRP() { return getToken(CmdParser.TOK_RMGRP, 0); }
		public TerminalNode TOK_MKUSR() { return getToken(CmdParser.TOK_MKUSR, 0); }
		public TerminalNode TOK_RMUSR() { return getToken(CmdParser.TOK_RMUSR, 0); }
		public TerminalNode TOK_MKFILE() { return getToken(CmdParser.TOK_MKFILE, 0); }
		public TerminalNode TOK_MKDIR() { return getToken(CmdParser.TOK_MKDIR, 0); }
		public TerminalNode TOK_EXEC() { return getToken(CmdParser.TOK_EXEC, 0); }
		public TerminalNode TOK_REP() { return getToken(CmdParser.TOK_REP, 0); }
		public TerminalNode TOK_PAUSE() { return getToken(CmdParser.TOK_PAUSE, 0); }
		public TerminalNode TOK_LOGOUT() { return getToken(CmdParser.TOK_LOGOUT, 0); }
		public ComandoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comando; }
	}

	public final ComandoContext comando() throws RecognitionException {
		ComandoContext _localctx = new ComandoContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_comando);
		try {
			setState(87);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case TOK_MKDISK:
				enterOuterAlt(_localctx, 1);
				{
				setState(27);
				match(TOK_MKDISK);
				setState(28);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewMkdisk(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_RMDISK:
				enterOuterAlt(_localctx, 2);
				{
				setState(31);
				match(TOK_RMDISK);
				setState(32);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewRmdisk(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_FDISK:
				enterOuterAlt(_localctx, 3);
				{
				setState(35);
				match(TOK_FDISK);
				setState(36);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewFdisk(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_MOUNT:
				enterOuterAlt(_localctx, 4);
				{
				setState(39);
				match(TOK_MOUNT);
				setState(40);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewMount(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_MKFS:
				enterOuterAlt(_localctx, 5);
				{
				setState(43);
				match(TOK_MKFS);
				setState(44);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewMkfs(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_LOGIN:
				enterOuterAlt(_localctx, 6);
				{
				setState(47);
				match(TOK_LOGIN);
				setState(48);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewLogin(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_MKGRP:
				enterOuterAlt(_localctx, 7);
				{
				setState(51);
				match(TOK_MKGRP);
				setState(52);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewMkgrp(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_RMGRP:
				enterOuterAlt(_localctx, 8);
				{
				setState(55);
				match(TOK_RMGRP);
				setState(56);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewRmgrp(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_MKUSR:
				enterOuterAlt(_localctx, 9);
				{
				setState(59);
				match(TOK_MKUSR);
				setState(60);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewMkusr(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_RMUSR:
				enterOuterAlt(_localctx, 10);
				{
				setState(63);
				match(TOK_RMUSR);
				setState(64);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewRmusr(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_MKFILE:
				enterOuterAlt(_localctx, 11);
				{
				setState(67);
				match(TOK_MKFILE);
				setState(68);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewMkfile(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_MKDIR:
				enterOuterAlt(_localctx, 12);
				{
				setState(71);
				match(TOK_MKDIR);
				setState(72);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewMkdir(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_EXEC:
				enterOuterAlt(_localctx, 13);
				{
				setState(75);
				match(TOK_EXEC);
				setState(76);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewExec(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_REP:
				enterOuterAlt(_localctx, 14);
				{
				setState(79);
				match(TOK_REP);
				setState(80);
				((ComandoContext)_localctx).param_list = param_list(0);

				                       _localctx.com = ast.NewRep(((ComandoContext)_localctx).param_list.lista)
				                
				}
				break;
			case TOK_PAUSE:
				enterOuterAlt(_localctx, 15);
				{
				setState(83);
				match(TOK_PAUSE);

				                       _localctx.com = ast.NewPause(true)
				                
				}
				break;
			case TOK_LOGOUT:
				enterOuterAlt(_localctx, 16);
				{
				setState(85);
				match(TOK_LOGOUT);

				                       _localctx.com = ast.NewLogout(true)
				                
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Param_listContext extends ParserRuleContext {
		public *arrayList.List lista;
		public Param_listContext AUXPLIST;
		public ParamContext param;
		public ParamContext param() {
			return getRuleContext(ParamContext.class,0);
		}
		public Param_listContext param_list() {
			return getRuleContext(Param_listContext.class,0);
		}
		public Param_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param_list; }
	}

	public final Param_listContext param_list() throws RecognitionException {
		return param_list(0);
	}

	private Param_listContext param_list(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Param_listContext _localctx = new Param_listContext(_ctx, _parentState);
		Param_listContext _prevctx = _localctx;
		int _startState = 6;
		enterRecursionRule(_localctx, 6, RULE_param_list, _p);
		_localctx.lista = arrayList.New()
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(90);
			((Param_listContext)_localctx).param = param();

			                   _localctx.lista.Add(((Param_listContext)_localctx).param.par)
			                
			}
			_ctx.stop = _input.LT(-1);
			setState(99);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new Param_listContext(_parentctx, _parentState);
					_localctx.AUXPLIST = _prevctx;
					_localctx.AUXPLIST = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_param_list);
					setState(93);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(94);
					((Param_listContext)_localctx).param = param();

					                             ((Param_listContext)_localctx).AUXPLIST.lista.Add(((Param_listContext)_localctx).param.par)
					                             _localctx.lista = ((Param_listContext)_localctx).AUXPLIST.lista
					                          
					}
					} 
				}
				setState(101);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ParamContext extends ParserRuleContext {
		public parametros.Parametro par;
		public Token TOK_NUMERO;
		public Token TOK_CAMINO;
		public Token TOK_CADENA;
		public Token TOK_PALABRA;
		public Token TOK_IDENTIFICADOR;
		public TerminalNode TOK_SIZE() { return getToken(CmdParser.TOK_SIZE, 0); }
		public TerminalNode TOK_IGUAL() { return getToken(CmdParser.TOK_IGUAL, 0); }
		public TerminalNode TOK_NUMERO() { return getToken(CmdParser.TOK_NUMERO, 0); }
		public TerminalNode TOK_PATH() { return getToken(CmdParser.TOK_PATH, 0); }
		public TerminalNode TOK_CAMINO() { return getToken(CmdParser.TOK_CAMINO, 0); }
		public TerminalNode TOK_CADENA() { return getToken(CmdParser.TOK_CADENA, 0); }
		public TerminalNode TOK_FIT() { return getToken(CmdParser.TOK_FIT, 0); }
		public TerminalNode TOK_FIRST() { return getToken(CmdParser.TOK_FIRST, 0); }
		public TerminalNode TOK_WORST() { return getToken(CmdParser.TOK_WORST, 0); }
		public TerminalNode TOK_BEST() { return getToken(CmdParser.TOK_BEST, 0); }
		public TerminalNode TOK_UNIT() { return getToken(CmdParser.TOK_UNIT, 0); }
		public TerminalNode TOK_BYTES() { return getToken(CmdParser.TOK_BYTES, 0); }
		public TerminalNode TOK_KB() { return getToken(CmdParser.TOK_KB, 0); }
		public TerminalNode TOK_MB() { return getToken(CmdParser.TOK_MB, 0); }
		public TerminalNode TOK_NAME() { return getToken(CmdParser.TOK_NAME, 0); }
		public TerminalNode TOK_PALABRA() { return getToken(CmdParser.TOK_PALABRA, 0); }
		public TerminalNode TOK_USUARIO() { return getToken(CmdParser.TOK_USUARIO, 0); }
		public TerminalNode TOK_GRP() { return getToken(CmdParser.TOK_GRP, 0); }
		public TerminalNode TOK_PASSWORD() { return getToken(CmdParser.TOK_PASSWORD, 0); }
		public TerminalNode TOK_PWD() { return getToken(CmdParser.TOK_PWD, 0); }
		public TerminalNode TOK_TYPE() { return getToken(CmdParser.TOK_TYPE, 0); }
		public TerminalNode TOK_PRIMARIA() { return getToken(CmdParser.TOK_PRIMARIA, 0); }
		public TerminalNode TOK_LOGICA() { return getToken(CmdParser.TOK_LOGICA, 0); }
		public TerminalNode TOK_EXTENDIDA() { return getToken(CmdParser.TOK_EXTENDIDA, 0); }
		public TerminalNode TOK_FAST() { return getToken(CmdParser.TOK_FAST, 0); }
		public TerminalNode TOK_FULL() { return getToken(CmdParser.TOK_FULL, 0); }
		public TerminalNode TOK_ID() { return getToken(CmdParser.TOK_ID, 0); }
		public TerminalNode TOK_IDENTIFICADOR() { return getToken(CmdParser.TOK_IDENTIFICADOR, 0); }
		public TerminalNode TOK_CONT() { return getToken(CmdParser.TOK_CONT, 0); }
		public TerminalNode TOK_RUTA() { return getToken(CmdParser.TOK_RUTA, 0); }
		public TerminalNode TOK_P() { return getToken(CmdParser.TOK_P, 0); }
		public TerminalNode TOK_R() { return getToken(CmdParser.TOK_R, 0); }
		public ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param; }
	}

	public final ParamContext param() throws RecognitionException {
		ParamContext _localctx = new ParamContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_param);
		try {
			setState(214);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(102);
				match(TOK_SIZE);
				setState(103);
				match(TOK_IGUAL);
				setState(104);
				((ParamContext)_localctx).TOK_NUMERO = match(TOK_NUMERO);

				              _localctx.par = parametros.NewParametro(ast.SIZE, (((ParamContext)_localctx).TOK_NUMERO!=null?((ParamContext)_localctx).TOK_NUMERO.getText():null))
				        
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(106);
				match(TOK_PATH);
				setState(107);
				match(TOK_IGUAL);
				setState(108);
				((ParamContext)_localctx).TOK_CAMINO = match(TOK_CAMINO);

				              _localctx.par = parametros.NewParametro(ast.PATH, (((ParamContext)_localctx).TOK_CAMINO!=null?((ParamContext)_localctx).TOK_CAMINO.getText():null))
				        
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(110);
				match(TOK_PATH);
				setState(111);
				match(TOK_IGUAL);
				setState(112);
				((ParamContext)_localctx).TOK_CADENA = match(TOK_CADENA);

				              str := strings.Trim((((ParamContext)_localctx).TOK_CADENA!=null?((ParamContext)_localctx).TOK_CADENA.getText():null), "\"")
				              _localctx.par = parametros.NewParametro(ast.PATH, str)
				        
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(114);
				match(TOK_FIT);
				setState(115);
				match(TOK_IGUAL);
				setState(116);
				match(TOK_FIRST);

				              _localctx.par = parametros.NewParametro(ast.FIT, ast.FF)
				        
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(118);
				match(TOK_FIT);
				setState(119);
				match(TOK_IGUAL);
				setState(120);
				match(TOK_WORST);

				              _localctx.par = parametros.NewParametro(ast.FIT, ast.WF)
				        
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(122);
				match(TOK_FIT);
				setState(123);
				match(TOK_IGUAL);
				setState(124);
				match(TOK_BEST);

				              _localctx.par = parametros.NewParametro(ast.FIT, ast.BF)
				        
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(126);
				match(TOK_UNIT);
				setState(127);
				match(TOK_IGUAL);
				setState(128);
				match(TOK_BYTES);

				              _localctx.par = parametros.NewParametro(ast.UNIT, ast.B)
				        
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(130);
				match(TOK_UNIT);
				setState(131);
				match(TOK_IGUAL);
				setState(132);
				match(TOK_KB);

				              _localctx.par = parametros.NewParametro(ast.UNIT, ast.K)
				        
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(134);
				match(TOK_UNIT);
				setState(135);
				match(TOK_IGUAL);
				setState(136);
				match(TOK_MB);

				              _localctx.par = parametros.NewParametro(ast.UNIT, ast.M)
				        
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(138);
				match(TOK_NAME);
				setState(139);
				match(TOK_IGUAL);
				setState(140);
				((ParamContext)_localctx).TOK_PALABRA = match(TOK_PALABRA);

				              _localctx.par = parametros.NewParametro(ast.NAME, (((ParamContext)_localctx).TOK_PALABRA!=null?((ParamContext)_localctx).TOK_PALABRA.getText():null))
				        
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(142);
				match(TOK_NAME);
				setState(143);
				match(TOK_IGUAL);
				setState(144);
				((ParamContext)_localctx).TOK_CADENA = match(TOK_CADENA);

				              str := strings.Trim((((ParamContext)_localctx).TOK_CADENA!=null?((ParamContext)_localctx).TOK_CADENA.getText():null), "\"")
				              _localctx.par = parametros.NewParametro(ast.NAME, str)
				        
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(146);
				match(TOK_USUARIO);
				setState(147);
				match(TOK_IGUAL);
				setState(148);
				((ParamContext)_localctx).TOK_PALABRA = match(TOK_PALABRA);

				              _localctx.par = parametros.NewParametro(ast.USUARIO, (((ParamContext)_localctx).TOK_PALABRA!=null?((ParamContext)_localctx).TOK_PALABRA.getText():null))
				        
				}
				break;
			case 13:
				enterOuterAlt(_localctx, 13);
				{
				setState(150);
				match(TOK_USUARIO);
				setState(151);
				match(TOK_IGUAL);
				setState(152);
				((ParamContext)_localctx).TOK_CADENA = match(TOK_CADENA);

				              str := strings.Trim((((ParamContext)_localctx).TOK_CADENA!=null?((ParamContext)_localctx).TOK_CADENA.getText():null), "\"")
				              _localctx.par = parametros.NewParametro(ast.USUARIO, str)
				        
				}
				break;
			case 14:
				enterOuterAlt(_localctx, 14);
				{
				setState(154);
				match(TOK_GRP);
				setState(155);
				match(TOK_IGUAL);
				setState(156);
				((ParamContext)_localctx).TOK_PALABRA = match(TOK_PALABRA);

				              _localctx.par = parametros.NewParametro(ast.GRP, (((ParamContext)_localctx).TOK_PALABRA!=null?((ParamContext)_localctx).TOK_PALABRA.getText():null))
				        
				}
				break;
			case 15:
				enterOuterAlt(_localctx, 15);
				{
				setState(158);
				match(TOK_GRP);
				setState(159);
				match(TOK_IGUAL);
				setState(160);
				((ParamContext)_localctx).TOK_CADENA = match(TOK_CADENA);

				              str := strings.Trim((((ParamContext)_localctx).TOK_CADENA!=null?((ParamContext)_localctx).TOK_CADENA.getText():null), "\"")
				              _localctx.par = parametros.NewParametro(ast.GRP, str)
				        
				}
				break;
			case 16:
				enterOuterAlt(_localctx, 16);
				{
				setState(162);
				match(TOK_PASSWORD);
				setState(163);
				match(TOK_IGUAL);
				setState(164);
				((ParamContext)_localctx).TOK_PALABRA = match(TOK_PALABRA);

				              _localctx.par = parametros.NewParametro(ast.PASSWORD, (((ParamContext)_localctx).TOK_PALABRA!=null?((ParamContext)_localctx).TOK_PALABRA.getText():null))
				        
				}
				break;
			case 17:
				enterOuterAlt(_localctx, 17);
				{
				setState(166);
				match(TOK_PWD);
				setState(167);
				match(TOK_IGUAL);
				setState(168);
				((ParamContext)_localctx).TOK_PALABRA = match(TOK_PALABRA);

				              _localctx.par = parametros.NewParametro(ast.PWD, (((ParamContext)_localctx).TOK_PALABRA!=null?((ParamContext)_localctx).TOK_PALABRA.getText():null))
				        
				}
				break;
			case 18:
				enterOuterAlt(_localctx, 18);
				{
				setState(170);
				match(TOK_TYPE);
				setState(171);
				match(TOK_IGUAL);
				setState(172);
				match(TOK_PRIMARIA);

				              _localctx.par = parametros.NewParametro(ast.TYPE, ast.PP)
				        
				}
				break;
			case 19:
				enterOuterAlt(_localctx, 19);
				{
				setState(174);
				match(TOK_TYPE);
				setState(175);
				match(TOK_IGUAL);
				setState(176);
				match(TOK_LOGICA);

				              _localctx.par = parametros.NewParametro(ast.TYPE, ast.L)
				        
				}
				break;
			case 20:
				enterOuterAlt(_localctx, 20);
				{
				setState(178);
				match(TOK_TYPE);
				setState(179);
				match(TOK_IGUAL);
				setState(180);
				match(TOK_EXTENDIDA);

				              _localctx.par = parametros.NewParametro(ast.TYPE, ast.E)
				        
				}
				break;
			case 21:
				enterOuterAlt(_localctx, 21);
				{
				setState(182);
				match(TOK_TYPE);
				setState(183);
				match(TOK_IGUAL);
				setState(184);
				match(TOK_FAST);

				              _localctx.par = parametros.NewParametro(ast.TYPE, ast.FAST)
				        
				}
				break;
			case 22:
				enterOuterAlt(_localctx, 22);
				{
				setState(186);
				match(TOK_TYPE);
				setState(187);
				match(TOK_IGUAL);
				setState(188);
				match(TOK_FULL);

				              _localctx.par = parametros.NewParametro(ast.TYPE, ast.FULL)
				        
				}
				break;
			case 23:
				enterOuterAlt(_localctx, 23);
				{
				setState(190);
				match(TOK_ID);
				setState(191);
				match(TOK_IGUAL);
				setState(192);
				((ParamContext)_localctx).TOK_IDENTIFICADOR = match(TOK_IDENTIFICADOR);

				              _localctx.par = parametros.NewParametro(ast.ID, (((ParamContext)_localctx).TOK_IDENTIFICADOR!=null?((ParamContext)_localctx).TOK_IDENTIFICADOR.getText():null))
				        
				}
				break;
			case 24:
				enterOuterAlt(_localctx, 24);
				{
				setState(194);
				match(TOK_CONT);
				setState(195);
				match(TOK_IGUAL);
				setState(196);
				((ParamContext)_localctx).TOK_CADENA = match(TOK_CADENA);

				              str := strings.Trim((((ParamContext)_localctx).TOK_CADENA!=null?((ParamContext)_localctx).TOK_CADENA.getText():null), "\"")
				              _localctx.par = parametros.NewParametro(ast.CONT, str)
				        
				}
				break;
			case 25:
				enterOuterAlt(_localctx, 25);
				{
				setState(198);
				match(TOK_CONT);
				setState(199);
				match(TOK_IGUAL);
				setState(200);
				((ParamContext)_localctx).TOK_CAMINO = match(TOK_CAMINO);

				              _localctx.par = parametros.NewParametro(ast.CONT, (((ParamContext)_localctx).TOK_CAMINO!=null?((ParamContext)_localctx).TOK_CAMINO.getText():null))
				        
				}
				break;
			case 26:
				enterOuterAlt(_localctx, 26);
				{
				setState(202);
				match(TOK_RUTA);
				setState(203);
				match(TOK_IGUAL);
				setState(204);
				((ParamContext)_localctx).TOK_CADENA = match(TOK_CADENA);

				              str := strings.Trim((((ParamContext)_localctx).TOK_CADENA!=null?((ParamContext)_localctx).TOK_CADENA.getText():null), "\"")
				              _localctx.par = parametros.NewParametro(ast.RUTA, str)
				        
				}
				break;
			case 27:
				enterOuterAlt(_localctx, 27);
				{
				setState(206);
				match(TOK_RUTA);
				setState(207);
				match(TOK_IGUAL);
				setState(208);
				((ParamContext)_localctx).TOK_CAMINO = match(TOK_CAMINO);

				              _localctx.par = parametros.NewParametro(ast.RUTA, (((ParamContext)_localctx).TOK_CAMINO!=null?((ParamContext)_localctx).TOK_CAMINO.getText():null))
				        
				}
				break;
			case 28:
				enterOuterAlt(_localctx, 28);
				{
				setState(210);
				match(TOK_P);

				              _localctx.par = parametros.NewParametro(ast.P, "true")
				        
				}
				break;
			case 29:
				enterOuterAlt(_localctx, 29);
				{
				setState(212);
				match(TOK_R);

				              _localctx.par = parametros.NewParametro(ast.R, "true")
				        
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 1:
			return comandList_sempred((ComandListContext)_localctx, predIndex);
		case 3:
			return param_list_sempred((Param_listContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean comandList_sempred(ComandListContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean param_list_sempred(Param_listContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3\64\u00db\4\2\t\2"+
		"\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\3\2\3\2\3\2\3\2\3\3\3\3\3\3\3\3\3\3\3"+
		"\3\3\3\3\3\7\3\31\n\3\f\3\16\3\34\13\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5"+
		"\4Z\n\4\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\7\5d\n\5\f\5\16\5g\13\5\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6"+
		"\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3\6\3"+
		"\6\3\6\3\6\3\6\3\6\3\6\5\6\u00d9\n\6\3\6\2\4\4\b\7\2\4\6\b\n\2\2\2\u0102"+
		"\2\f\3\2\2\2\4\20\3\2\2\2\6Y\3\2\2\2\b[\3\2\2\2\n\u00d8\3\2\2\2\f\r\5"+
		"\4\3\2\r\16\b\2\1\2\16\17\7\2\2\3\17\3\3\2\2\2\20\21\b\3\1\2\21\22\5\6"+
		"\4\2\22\23\b\3\1\2\23\32\3\2\2\2\24\25\f\4\2\2\25\26\5\6\4\2\26\27\b\3"+
		"\1\2\27\31\3\2\2\2\30\24\3\2\2\2\31\34\3\2\2\2\32\30\3\2\2\2\32\33\3\2"+
		"\2\2\33\5\3\2\2\2\34\32\3\2\2\2\35\36\7\4\2\2\36\37\5\b\5\2\37 \b\4\1"+
		"\2 Z\3\2\2\2!\"\7\5\2\2\"#\5\b\5\2#$\b\4\1\2$Z\3\2\2\2%&\7\6\2\2&\'\5"+
		"\b\5\2\'(\b\4\1\2(Z\3\2\2\2)*\7\7\2\2*+\5\b\5\2+,\b\4\1\2,Z\3\2\2\2-."+
		"\7\b\2\2./\5\b\5\2/\60\b\4\1\2\60Z\3\2\2\2\61\62\7\t\2\2\62\63\5\b\5\2"+
		"\63\64\b\4\1\2\64Z\3\2\2\2\65\66\7\13\2\2\66\67\5\b\5\2\678\b\4\1\28Z"+
		"\3\2\2\29:\7\f\2\2:;\5\b\5\2;<\b\4\1\2<Z\3\2\2\2=>\7\r\2\2>?\5\b\5\2?"+
		"@\b\4\1\2@Z\3\2\2\2AB\7\16\2\2BC\5\b\5\2CD\b\4\1\2DZ\3\2\2\2EF\7\17\2"+
		"\2FG\5\b\5\2GH\b\4\1\2HZ\3\2\2\2IJ\7\20\2\2JK\5\b\5\2KL\b\4\1\2LZ\3\2"+
		"\2\2MN\7\21\2\2NO\5\b\5\2OP\b\4\1\2PZ\3\2\2\2QR\7\23\2\2RS\5\b\5\2ST\b"+
		"\4\1\2TZ\3\2\2\2UV\7\22\2\2VZ\b\4\1\2WX\7\n\2\2XZ\b\4\1\2Y\35\3\2\2\2"+
		"Y!\3\2\2\2Y%\3\2\2\2Y)\3\2\2\2Y-\3\2\2\2Y\61\3\2\2\2Y\65\3\2\2\2Y9\3\2"+
		"\2\2Y=\3\2\2\2YA\3\2\2\2YE\3\2\2\2YI\3\2\2\2YM\3\2\2\2YQ\3\2\2\2YU\3\2"+
		"\2\2YW\3\2\2\2Z\7\3\2\2\2[\\\b\5\1\2\\]\5\n\6\2]^\b\5\1\2^e\3\2\2\2_`"+
		"\f\4\2\2`a\5\n\6\2ab\b\5\1\2bd\3\2\2\2c_\3\2\2\2dg\3\2\2\2ec\3\2\2\2e"+
		"f\3\2\2\2f\t\3\2\2\2ge\3\2\2\2hi\7\26\2\2ij\7\3\2\2jk\7/\2\2k\u00d9\b"+
		"\6\1\2lm\7\24\2\2mn\7\3\2\2no\7\61\2\2o\u00d9\b\6\1\2pq\7\24\2\2qr\7\3"+
		"\2\2rs\7.\2\2s\u00d9\b\6\1\2tu\7\25\2\2uv\7\3\2\2vw\7)\2\2w\u00d9\b\6"+
		"\1\2xy\7\25\2\2yz\7\3\2\2z{\7*\2\2{\u00d9\b\6\1\2|}\7\25\2\2}~\7\3\2\2"+
		"~\177\7+\2\2\177\u00d9\b\6\1\2\u0080\u0081\7\27\2\2\u0081\u0082\7\3\2"+
		"\2\u0082\u0083\7%\2\2\u0083\u00d9\b\6\1\2\u0084\u0085\7\27\2\2\u0085\u0086"+
		"\7\3\2\2\u0086\u0087\7#\2\2\u0087\u00d9\b\6\1\2\u0088\u0089\7\27\2\2\u0089"+
		"\u008a\7\3\2\2\u008a\u008b\7$\2\2\u008b\u00d9\b\6\1\2\u008c\u008d\7\30"+
		"\2\2\u008d\u008e\7\3\2\2\u008e\u008f\7\62\2\2\u008f\u00d9\b\6\1\2\u0090"+
		"\u0091\7\30\2\2\u0091\u0092\7\3\2\2\u0092\u0093\7.\2\2\u0093\u00d9\b\6"+
		"\1\2\u0094\u0095\7\33\2\2\u0095\u0096\7\3\2\2\u0096\u0097\7\62\2\2\u0097"+
		"\u00d9\b\6\1\2\u0098\u0099\7\33\2\2\u0099\u009a\7\3\2\2\u009a\u009b\7"+
		".\2\2\u009b\u00d9\b\6\1\2\u009c\u009d\7\37\2\2\u009d\u009e\7\3\2\2\u009e"+
		"\u009f\7\62\2\2\u009f\u00d9\b\6\1\2\u00a0\u00a1\7\37\2\2\u00a1\u00a2\7"+
		"\3\2\2\u00a2\u00a3\7.\2\2\u00a3\u00d9\b\6\1\2\u00a4\u00a5\7\34\2\2\u00a5"+
		"\u00a6\7\3\2\2\u00a6\u00a7\7\62\2\2\u00a7\u00d9\b\6\1\2\u00a8\u00a9\7"+
		"\35\2\2\u00a9\u00aa\7\3\2\2\u00aa\u00ab\7\62\2\2\u00ab\u00d9\b\6\1\2\u00ac"+
		"\u00ad\7\31\2\2\u00ad\u00ae\7\3\2\2\u00ae\u00af\7&\2\2\u00af\u00d9\b\6"+
		"\1\2\u00b0\u00b1\7\31\2\2\u00b1\u00b2\7\3\2\2\u00b2\u00b3\7(\2\2\u00b3"+
		"\u00d9\b\6\1\2\u00b4\u00b5\7\31\2\2\u00b5\u00b6\7\3\2\2\u00b6\u00b7\7"+
		"\'\2\2\u00b7\u00d9\b\6\1\2\u00b8\u00b9\7\31\2\2\u00b9\u00ba\7\3\2\2\u00ba"+
		"\u00bb\7,\2\2\u00bb\u00d9\b\6\1\2\u00bc\u00bd\7\31\2\2\u00bd\u00be\7\3"+
		"\2\2\u00be\u00bf\7-\2\2\u00bf\u00d9\b\6\1\2\u00c0\u00c1\7\32\2\2\u00c1"+
		"\u00c2\7\3\2\2\u00c2\u00c3\7\60\2\2\u00c3\u00d9\b\6\1\2\u00c4\u00c5\7"+
		"\36\2\2\u00c5\u00c6\7\3\2\2\u00c6\u00c7\7.\2\2\u00c7\u00d9\b\6\1\2\u00c8"+
		"\u00c9\7\36\2\2\u00c9\u00ca\7\3\2\2\u00ca\u00cb\7\61\2\2\u00cb\u00d9\b"+
		"\6\1\2\u00cc\u00cd\7 \2\2\u00cd\u00ce\7\3\2\2\u00ce\u00cf\7.\2\2\u00cf"+
		"\u00d9\b\6\1\2\u00d0\u00d1\7 \2\2\u00d1\u00d2\7\3\2\2\u00d2\u00d3\7\61"+
		"\2\2\u00d3\u00d9\b\6\1\2\u00d4\u00d5\7\"\2\2\u00d5\u00d9\b\6\1\2\u00d6"+
		"\u00d7\7!\2\2\u00d7\u00d9\b\6\1\2\u00d8h\3\2\2\2\u00d8l\3\2\2\2\u00d8"+
		"p\3\2\2\2\u00d8t\3\2\2\2\u00d8x\3\2\2\2\u00d8|\3\2\2\2\u00d8\u0080\3\2"+
		"\2\2\u00d8\u0084\3\2\2\2\u00d8\u0088\3\2\2\2\u00d8\u008c\3\2\2\2\u00d8"+
		"\u0090\3\2\2\2\u00d8\u0094\3\2\2\2\u00d8\u0098\3\2\2\2\u00d8\u009c\3\2"+
		"\2\2\u00d8\u00a0\3\2\2\2\u00d8\u00a4\3\2\2\2\u00d8\u00a8\3\2\2\2\u00d8"+
		"\u00ac\3\2\2\2\u00d8\u00b0\3\2\2\2\u00d8\u00b4\3\2\2\2\u00d8\u00b8\3\2"+
		"\2\2\u00d8\u00bc\3\2\2\2\u00d8\u00c0\3\2\2\2\u00d8\u00c4\3\2\2\2\u00d8"+
		"\u00c8\3\2\2\2\u00d8\u00cc\3\2\2\2\u00d8\u00d0\3\2\2\2\u00d8\u00d4\3\2"+
		"\2\2\u00d8\u00d6\3\2\2\2\u00d9\13\3\2\2\2\6\32Ye\u00d8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}
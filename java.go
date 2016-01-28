/*

import java.io.ByteArrayInputStream;
import java.io.ByteArrayOutputStream;
import java.io.InputStream;
import java.io.OutputStream;
import java.security.Key;
import java.security.KeyFactory;
import java.security.KeyPair;
import java.security.KeyPairGenerator;
import java.security.PrivateKey;
import java.security.PublicKey;
import java.security.SecureRandom;
import java.security.spec.PKCS8EncodedKeySpec;
import java.security.spec.X509EncodedKeySpec;

import javax.crypto.Cipher;

import org.apache.commons.codec.binary.Base64;
import org.apache.commons.io.IOUtils;

public class RSASecurity {
	
	private PublicKey pubKey;
	private PrivateKey priKey;
	private KeySize size;
	
	public RSASecurity(String pub, String pri, KeySize size) throws PubkeyException, PrikeyException {
		this.size = size;
		this.pubKey = getPublicKey(pub);
		this.priKey = getPrivateKey(pri);
	}
	
	public class PubkeyException extends Exception {
		private static final long serialVersionUID = 2416296748628947909L;
		public PubkeyException(Throwable cause){
			super(cause);
			
		}
	}
	
	public class PrikeyException extends Exception {
		private static final long serialVersionUID = -233684887155150306L;
		public PrikeyException(Throwable cause){
			super(cause);
		}
	}
	
	public enum Mode {
		PUBKEY_ENCRYPT, //公钥加密
		PUBKEY_DECRYPT, //公钥解密
		PRIKEY_ENCRYPT, //私钥加密
		PRIKEY_DECRYPT;  //私钥解密
	}
	
	public enum KeySize{
		SIZE_1024(1024), SIZE_2048(2048), SIZE_768(768);
		private int _value;
		private KeySize(int value){
			_value = value;
		}
		public int value(){
			return _value;
		}
	}
	
	public String exec(String data, Mode mode) throws Exception{
		if(data == null){
			return null;
		}
		if(mode == Mode.PUBKEY_ENCRYPT || mode == Mode.PRIKEY_ENCRYPT){
			return Base64.encodeBase64String(exec(data.getBytes("UTF-8"), mode));
		}else{
			return new String(exec(Base64.decodeBase64(data), mode), "UTF-8");
		}
	}
	
	public byte[] exec(byte[] data, Mode mode) throws Exception{
		if(data == null){
			return null;
		}
		ByteArrayOutputStream baos = null;
		ByteArrayInputStream bais = null;
		try{
			baos = new ByteArrayOutputStream();
			bais = new ByteArrayInputStream(data);
			exec(bais, baos, mode);
			return baos.toByteArray();
		}finally{
			IOUtils.closeQuietly(baos);
			IOUtils.closeQuietly(bais);
		}
	}
	
	
	public void exec(InputStream in, OutputStream os, Mode mode) throws Exception{
		Cipher cipher = Cipher.getInstance("RSA/ECB/PKCS1Padding");
		switch (mode) {
		case PUBKEY_ENCRYPT:
			cipher.init(Cipher.ENCRYPT_MODE, pubKey);
			encrypt(cipher, in, os);
			break;
		case PUBKEY_DECRYPT:
			cipher.init(Cipher.DECRYPT_MODE, pubKey);
			decrypt(cipher, in, os);
			break;
		case PRIKEY_ENCRYPT:
			cipher.init(Cipher.ENCRYPT_MODE, priKey);
			encrypt(cipher, in, os);
			break;
		case PRIKEY_DECRYPT:
			cipher.init(Cipher.DECRYPT_MODE, priKey);
			decrypt(cipher, in, os);
			break;
		}
	}
	
	//加密
	private void encrypt(Cipher cipher, InputStream in, OutputStream os) throws Exception{
		byte[] buf = new byte[size.value()/8-11];
		int len = 0;
		while ((len = in.read(buf)) != -1) {
			os.write(cipher.doFinal(buf, 0, len));
		}
	}
	//解密
	private void decrypt(Cipher cipher, InputStream in, OutputStream os) throws Exception{
		byte[] buf = new byte[size.value()/8];
		while (in.read(buf) != -1) {
			os.write(cipher.doFinal(buf));
		}
	}

	//获取公钥Cipher
	private PublicKey getPublicKey(String pubString) throws PubkeyException  {
		try{
			byte[] keyBytes = Base64.decodeBase64(pubString);
			X509EncodedKeySpec keySpec = new X509EncodedKeySpec(keyBytes);
			KeyFactory keyFactory = KeyFactory.getInstance("RSA");
			PublicKey publicKey = keyFactory.generatePublic(keySpec);
			return publicKey;
		}catch(Exception e){
			throw new PubkeyException(e);
		}
		
	}
	//获取私钥
	private PrivateKey getPrivateKey(String priString) throws PrikeyException {
		try{
			byte[] keyBytes = Base64.decodeBase64(priString);
			PKCS8EncodedKeySpec keySpec = new PKCS8EncodedKeySpec(keyBytes);
			KeyFactory keyFactory = KeyFactory.getInstance("RSA");
			PrivateKey privateKey = keyFactory.generatePrivate(keySpec);
			return privateKey;
		}catch(Exception e){
			throw new PrikeyException(e);
		}
	}
	// 生成公私钥
	public static void generateKeyPair(KeySize size) throws Exception {
		SecureRandom secureRandom = new SecureRandom();
		KeyPairGenerator keyPairGenerator = KeyPairGenerator.getInstance("RSA");
		keyPairGenerator.initialize(size.value(), secureRandom);
		keyPairGenerator.initialize(size.value());
		KeyPair keyPair = keyPairGenerator.generateKeyPair();
		Key publicKey = keyPair.getPublic();
		Key privateKey = keyPair.getPrivate();
		System.out.println("publickey:");
		char[] pubKey = Base64.encodeBase64String(publicKey.getEncoded()).toCharArray();
		for (int i=0; i<pubKey.length; i++){
			System.out.print(pubKey[i]);
			//if (i % 64 == 63){
				//System.out.println();
			//}
		}
		System.out.println();
		System.out.println("privatekey:");
		char[] priKey = Base64.encodeBase64String(privateKey.getEncoded()).toCharArray();
		for (int i=0; i<priKey.length; i++){
			System.out.print(priKey[i]);
			//if (i % 64 == 63){
				//System.out.println();
			//}
		}
		System.out.println();
	}
}
*/
